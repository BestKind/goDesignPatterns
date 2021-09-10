package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	r1 := new(Example)
	r1.Content = "This is 1"
	r2 := r1.Clone()
	r2.Content = "This is clone"

	fmt.Println(r1)
	fmt.Println(r2)
}
