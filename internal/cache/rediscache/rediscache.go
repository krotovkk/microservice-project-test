package rediscache

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type RedisCache struct {
	client *redis.Client

	productCache ports.ProductCache
	cartCache    ports.CartCache
}

func NewRedisCache(client *redis.Client) *RedisCache {
	cache := &RedisCache{
		client: client,
	}

	cache.productCache = NewProductRedisStore(cache)
	cache.cartCache = NewCartRedisCache(cache)

	return cache
}

func (s *RedisCache) Product() ports.ProductCache {
	return s.productCache
}

func (s *RedisCache) Cart() ports.CartCache {
	return s.cartCache
}

func (s *RedisCache) removeByKeys(pattern string) error {
	res := s.client.Keys(pattern)

	if res.Err() != nil {
		return res.Err()
	}

	dres := s.client.Del(res.Val()...)

	if dres.Err() != nil {
		return dres.Err()
	}

	return nil
}

func (s *RedisCache) sendProductsToChanel(products ProductsData, chanel string) error {
	data, err := json.Marshal(products)

	if err != nil {
		return err
	}

	res := s.client.Publish(chanel, data)

	if res.Err() != nil {
		return res.Err()
	}

	return nil
}
