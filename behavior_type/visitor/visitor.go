package visitor

import "fmt"

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForTriangle(*triangle)
}

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) getType() string {
	return "Square"
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

type circle struct {
	radius int
}

func (c *circle) getType() string {
	return "Square"
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

type triangle struct {
	h int
	b int
}

func (t *triangle) getType() string {
	return "Square"
}

func (t *triangle) accept(v visitor) {
	v.visitForTriangle(t)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("calculating area for square")
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("calculating area for circle")
}

func (a *areaCalculator) visitForTriangle(t *triangle) {
	fmt.Println("calculating area for triangle")
}
