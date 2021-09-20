package main

import "fmt"

func main() {
	x := Add(4, 5)
	fmt.Println(x)
}
func Add(x, y int) int {
	return x + y
}
