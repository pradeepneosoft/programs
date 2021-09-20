package main

import "fmt"

type Person interface {
	Foo()
}
type P1 struct {
	Name string
}
type P2 struct {
	Age int
}

func (p P1) Foo() {
	fmt.Println(p.Name)
}
func (r P2) Foo() {
	fmt.Println(r.Age)
}

func Display(per Person) {
	per.Foo()
}
func main() {
	p := P1{Name: "jlkdjskds"}
	r := P2{Age: 34}
	// var person Person
	// person = p
	Display(p)
	Display(r)
}
