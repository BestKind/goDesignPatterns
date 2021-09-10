package factory

import (
	"fmt"
	"testing"
)

func TestRedisCache(t *testing.T) {
	var redisCacheFactory CacheFactory
	redisCacheFactory = RedisCacheFactory{}

	redisCache := redisCacheFactory.Create()
	redisCache.Set("k1", "v1")

	value := redisCache.Get("k1")
	fmt.Println(value)
}
