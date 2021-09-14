package visitor

import "testing"

func TestVisitor(t *testing.T) {
	s := &square{side:2}
	c := &circle{radius:3}
	tr := &triangle{b:4, h:3}

	areaCalculator := &areaCalculator{}

	s.accept(areaCalculator)
	c.accept(areaCalculator)
	tr.accept(areaCalculator)
}
