package single_factory

import (
	"fmt"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	product1 := &Product1{}
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := &Product2{}
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}

func TestProductFactory_Create(t *testing.T) {
	pf := productFactory{}
	p1 := pf.Create(p1)
	p1.SetName("p1")
	fmt.Println(p1.GetName())

	p2 := pf.Create(p2)
	p2.SetName("p2")
	fmt.Println(p2.GetName())
}
