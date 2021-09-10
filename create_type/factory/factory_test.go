package factory

import (
	"fmt"
	"testing"
)

func TestProduct1(t *testing.T) {
	var productFactory1 ProductFactory

	productFactory1 = &Product1Factory{}
	p1 := productFactory1.Create()
	p1.SetName("p1")

	name := p1.GetName()
	fmt.Println(name)
}
