package main

import "fmt"

func main() {
	x := factorial(5)
	fmt.Println(x)
	y := withloop(7)
	fmt.Println(y)
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)

}
func withloop(x int) int {
	total := 1
	for ; x > 0; x-- {
		total = x * total
	}
	return total

}
