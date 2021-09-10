# 简单工厂模式
简单工厂模式（`Simple Factory Pattern`）：定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类。因为在简单工厂模式中用于创建实例的方法是静态（`static`）方法，因此简单工厂模式又被称为静态工厂方法（`Static Factory Method`）模式，它属于类创建型模式。

简单工厂模式的要点在于：当你需要什么，只需要传入一个正确的参数，就可以获取你所需要的对象，而无须知道其创建细节。

以上定义来自于网络

代码示例
```
package main

// 实现一个抽象的产品
type Product interface {
	SetName(name string)
	GetName() string
}

// 实现具体的产品：产品1
type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return "产品1的 name 为" + p1.name
}

// 实现具体的产品：产品2
type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return "产品2的 name 为" + p2.name
}

type productType int

const (
	p1 productType = iota
	p2
)

// 实现简单工厂类
type productFactory struct {
}

func (pf productFactory) Create(productType productType) Product {
	if p1 == productType {
		return &Product1{}
	}

	if p2 == productType {
		return &Product2{}
	}
	return nil
}
```

简单实现了一个 `cache` 的工厂模式
