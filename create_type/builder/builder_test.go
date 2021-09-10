package builder

import "testing"

func TestBuilder(t *testing.T) {
	director := new(Director)
	builder := new(CarBuilder)
	director.SetBuilder(builder)
	car := director.Generate()
	car.Show()
}
