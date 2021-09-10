package singleton

import (
	"sync"
)

type singletonLazy struct {
	name string
}

var instance *singletonLazy

func GetInstance() *singletonLazy {
	if instance == nil {
		instance = new(singletonLazy)
		instance.name = "懒汉单例"
	}
	return instance
}

type singletonFull struct {
	name string
}

var instance2 *singletonFull

func init() {
	instance2 = new(singletonFull)
	instance2.name = "饿汉单例"
}

func GetInstance2() *singletonFull {
	return instance2
}

var instance1 *singletonLazy
var once sync.Once

func GetInstance1() *singletonLazy {
	once.Do(func() {
		if instance1 == nil {
			instance1 = new(singletonLazy)
			instance1.name = "double check"
		}
	})
	return instance1
}
