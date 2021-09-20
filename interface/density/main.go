package main

import "fmt"

type Dense interface {
	Density() int
}

type Rock struct {
}
type Geode struct {
}

func (r Rock) Density() int {
	return 100
}
func (g Geode) Density() int {
	return 120
}
func isitDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

func main() {
	a := Rock{}
	b := Geode{}
	fmt.Println(isitDenser(a, b))
}
