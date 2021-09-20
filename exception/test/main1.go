package main

import (
	"fmt"
)

func main() {
	execPanic()
	fmt.Println("main print")
}
func execPanic() {
	defer recovry()
	panic("panic called ")
	fmt.Println("panic func")
}
func recovry() {
	if msg := recover(); msg != nil {
		fmt.Println(msg)
	}
	fmt.Println("recovry function")

}
