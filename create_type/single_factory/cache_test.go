package single_factory

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cf := &CacheFactory{}
	redis, err := cf.Create(redis)
	if nil != err {
		t.Error(err)
	}
	redis.Set("k1", "v1")
	fmt.Println(redis.Get("k1"))

	mem, err := cf.Create(mem)
	if nil != err {
		t.Error(err)
	}
	mem.Set("k1", "v1")
	fmt.Println(mem.Get("k1"))
}
