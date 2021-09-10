package single_factory

import "errors"

// 定义一个 Cache 接口 作为父类
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

// 实现具体的 Cache：RedisCache
type RedisCache struct {
	data map[string]string
}

func (redis *RedisCache) Set(key, value string) {
	redis.data[key] = value
}

func (redis *RedisCache) Get(key string) string {
	return "redis:" + redis.data[key]
}

// 实现具体的 Cache：MemCache
type MemCache struct {
	data map[string]string
}

func (mem *MemCache) Set(key, value string) {
	mem.data[key] = value
}

func (mem *MemCache) Get(key string) string {
	return "mem:" + mem.data[key]
}

type CacheType int

const (
	redis CacheType = iota
	mem
)

// 实现 Cache 的简单工厂
type CacheFactory struct {
}

func (factory *CacheFactory) Create(cacheType CacheType) (Cache, error) {
	if redis == cacheType {
		return &RedisCache{data: map[string]string{}}, nil
	}

	if mem == cacheType {
		return &MemCache{data: map[string]string{}}, nil
	}

	return nil, errors.New("Error cache type! ")
}
