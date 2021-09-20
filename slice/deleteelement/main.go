package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(x)
	x = append(x[:3], x[4:]...)
	fmt.Println(x)
}
