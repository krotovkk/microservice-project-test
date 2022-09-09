package rediscache

import (
	"context"
	"expvar"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/common"
	"gitlab.ozon.dev/krotovkk/homework/internal/counters"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

const cartProductsExpiration = time.Hour * 24

type CartRedisCache struct {
	missGet *counters.Counter
	hitGet  *counters.Counter

	*RedisCache
}

func NewCartRedisCache(cache *RedisCache) ports.CartCache {
	redisCache := &CartRedisCache{
		RedisCache: cache,
		missGet:    &counters.Counter{M: &sync.RWMutex{}},
		hitGet:     &counters.Counter{M: &sync.RWMutex{}},
	}

	expvar.Publish("cart cache miss get", redisCache.missGet)
	expvar.Publish("cart cache hit get", redisCache.hitGet)

	return redisCache
}

func (c *CartRedisCache) GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error) {
	res := c.client.Get(fmt.Sprintf("cart_products:%d", id))

	if res.Err() != nil {
		c.missGet.Inc()
		return nil, res.Err()
	}
	c.hitGet.Inc()

	var products ProductsData

	err := res.Scan(&products)

	if err != nil {
		return nil, err
	}

	err = c.SendProductsToChanel(products)

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "get cart products from cache"}).Error(err)
	}

	return products, nil
}

func (c *CartRedisCache) SetCartProducts(ctx context.Context, id int64, products []*model.Product) error {
	res := c.client.Set(fmt.Sprintf(fmt.Sprintf("cart_products:%d", id)), ProductsData(products), cartProductsExpiration)

	if res.Err() != nil {
		c.missGet.Inc()
		return res.Err()
	}
	c.hitGet.Inc()

	err := c.SendProductsToChanel(products)

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "set cart products from cache"}).Error(err)
	}

	return nil
}

func (c *CartRedisCache) ClearGetCartProducts(id int64) error {
	return c.removeByKeys(fmt.Sprintf("cart_products:%d", id))
}

func (c *CartRedisCache) ClearAllCartProducts() error {
	return c.removeByKeys("cart_products:*")
}

func (c *CartRedisCache) SendProductsToChanel(data []*model.Product) error {
	return c.sendProductsToChanel(data, common.CartProductsChanel)
}
