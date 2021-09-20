package main

import (
	"fmt"
	"math"
)

type geometry interface {
	perim() float64
	area() float64
}
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rect) area() float64 {
	return r.width * r.height
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (r rect) perim() float64 {
	return 2 * (r.width * r.height)
}
func Measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rect{height: 10, width: 20}
	c := circle{radius: 30}

	Measure(r)
	Measure(c)
}
