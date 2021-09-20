package main

import "fmt"

var Stack []int

func Pop() {
	if len(Stack) > 0 {
		Stack = Stack[:len(Stack)-1]
	} else {
		panic("Stack empty")
	}

}
func Push(element int) {
	Stack = append(Stack, element)
}
func main() {
	defer recovery()
	fmt.Println(Stack)
	Push(5)
	Push(50)
	Push(53)
	fmt.Println(Stack)
	Pop()
	Pop()
	Pop()
	Pop()
	fmt.Println(Stack)

}
func recovery() {
	if msg := recover(); msg != nil {
		fmt.Println(msg)
	}
}
