package rediscache

import (
	"context"
	"encoding/json"
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

const productsExpiration = time.Hour * 24

type ProductsData []*model.Product

func (p ProductsData) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *ProductsData) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}

type ProductRedisCache struct {
	missGet *counters.Counter
	hitGet  *counters.Counter

	*RedisCache
}

func NewProductRedisStore(cache *RedisCache) ports.ProductCache {
	redisCache := &ProductRedisCache{
		RedisCache: cache,
		missGet:    &counters.Counter{M: &sync.RWMutex{}},
		hitGet:     &counters.Counter{M: &sync.RWMutex{}},
	}

	expvar.Publish("product cache miss get", redisCache.missGet)
	expvar.Publish("product cache hit get", redisCache.hitGet)

	return redisCache
}

func (ps *ProductRedisCache) GetAllProducts(ctx context.Context, limit uint64, offset uint64) ([]*model.Product, error) {
	res := ps.client.Get(fmt.Sprintf("products:%d:%d", limit, offset))

	if res.Err() != nil {
		ps.missGet.Inc()
		return nil, res.Err()
	}
	ps.hitGet.Inc()

	var products ProductsData

	err := res.Scan(&products)

	if err != nil {
		return nil, err
	}

	err = ps.SendProductsToChanel(products)

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "set all products from cache"}).Error(err)
	}

	return products, nil
}

func (ps *ProductRedisCache) SetAllProducts(ctx context.Context, limit uint64, offset uint64, products []*model.Product) error {
	res := ps.client.Set(fmt.Sprintf("products:%d:%d", limit, offset), ProductsData(products), productsExpiration)

	if res.Err() != nil {
		ps.missGet.Inc()
		return res.Err()
	}
	ps.hitGet.Inc()

	err := ps.SendProductsToChanel(products)

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "set all products from cache"}).Error(err)
	}

	return nil
}

func (ps *ProductRedisCache) ClearGetAllProducts() error {
	return ps.removeByKeys("products:*")
}

func (ps *ProductRedisCache) SendProductsToChanel(data []*model.Product) error {
	return ps.sendProductsToChanel(data, common.ProductsListChanel)
}
