# 单例模式
单例模式要实现的效果就是，对于应用单例模式的类，整个程序中只存在一个实例化对象

有几种方式:
 - 懒汉模式
 - 饿汉模式
 - 双重检查锁机制

## 懒汉模式
代码示例
```go
package main

import "fmt"

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

func main() {
    lazy1 := GetInstance()
	fmt.Println(lazy1.name)
	lazy1.name = "lazy"

	lazy2 := GetInstance()
	fmt.Println(lazy2.name)
}
```

## 饿汉模式
代码示例
```go
package main

import "fmt"

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

func main() {
    full1 := GetInstance2()
    fmt.Println(full1.name)
    full1.name = "full"
    
    full2 := GetInstance2()
    fmt.Println(full2.name)
}
```

## 双重检查机制
代码示例
```go
package main

import (
    "fmt"
    "sync"
)

type singletonLazy struct {
	name string
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

func main() {
    d1 := GetInstance1()
    fmt.Println(d1.name)
    d1.name = "ddd"

    d2 := GetInstance1()
    fmt.Println(d2.name)
}
```
