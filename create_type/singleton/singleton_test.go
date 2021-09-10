package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	lazy1 := GetInstance()
	fmt.Println(lazy1.name)
	lazy1.name = "lazy"

	lazy2 := GetInstance()
	fmt.Println(lazy2.name)

	d1 := GetInstance1()
	fmt.Println(d1.name)
	d1.name = "ddd"

	d2 := GetInstance1()
	fmt.Println(d2.name)

	full1 := GetInstance2()
	fmt.Println(full1.name)
	full1.name = "full"

	full2 := GetInstance2()
	fmt.Println(full2.name)
}
