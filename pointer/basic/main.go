package main

import "fmt"

func main() {
	var a int
	var b *int
	a = 5
	b = &a
	fmt.Println(a)
	fmt.Println(*b)
}
