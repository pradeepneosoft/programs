package main

import (
	"fmt"
	"math"
)

type Sizer interface {
	Area() float64
}
type Shaper interface {
	Sizer
	Fun
	fmt.Stringer
}
type Fun interface {
	show()
}

type Circle struct {
	Radius float64
}
type Square struct {
	Width  float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Height * s.Width
}
func (c Circle) Area() float64 {
	return 2 * math.Pi * c.Radius
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)

}
func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)

}
func (s Square) show() {
	fmt.Println("fun Squarew")

}
func (c Circle) show() {
	fmt.Println("fun Circle")

}
func main() {
	s := Square{Width: 30, Height: 40}
	c := Circle{Radius: 14}
	PrintArea(s)
	PrintArea(c)

	l := Less(s, c)
	fmt.Println(l)
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() > s2.Area() {
		return s1
	}
	return s2

}
func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
	s.show()
}
