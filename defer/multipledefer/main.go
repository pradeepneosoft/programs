package main

import "fmt"

func main() {
	defer fmt.Println("Defer in main")
	fmt.Println("Stat main")
	f1()
	fmt.Println("Finish main")
}

func f1() {
	defer un(trim("f1"))
	fmt.Println("Start f1")
	f2()
	fmt.Println("Finish f1")
}

func f2() {
	defer un(trim("f2"))
	fmt.Println("Start f2")
	fmt.Println("Finish f2")
}
func trim(s string) string {
	fmt.Println("in  trim ", s)
	return s
}
func un(s string) {
	fmt.Println("Un", s)
}
