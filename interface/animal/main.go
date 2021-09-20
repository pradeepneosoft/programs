package main

import "fmt"

type animal interface {
	Speak() string
}

type Dog struct {
}
type Cat struct {
}
type Llama struct {
}
type JavaProgramer struct {
}

func (d Dog) Speak() string {
	return "bao"
}
func (c Cat) Speak() string {
	return "Meao"
}
func (l Llama) Speak() string {
	return "????"
}
func (j JavaProgramer) Speak() string {
	return "tak"
}
func main() {
	jaanwar := []animal{Dog{}, Cat{}, Llama{}, JavaProgramer{}}

	for _, j := range jaanwar {
		fmt.Println(j.Speak())
	}

}
