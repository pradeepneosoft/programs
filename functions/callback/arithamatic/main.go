package main

import "fmt"

func main() {
	math(add, 5, 6)
	math(multi, 5, 6)

}

func add(x, y int) {
	fmt.Println(x + y)
}
func multi(x, y int) {
	fmt.Println(x * y)

}
func math(f func(x, y int), x int, y int) {
	f(x, y)

}
