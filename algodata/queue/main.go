package main

import "fmt"

var Queue []int

func Enque(element int) {
	Queue = append(Queue, element)
}
func Deque() {
	if len(Queue) > 0 {
		Queue = Queue[1:]
	} else {
		panic("Queue Empty")

	}

}
func main() {
	fmt.Println(Queue)
	Enque(10)
	Enque(13)

	Enque(14)

	Enque(15)

	fmt.Println(Queue)
	Deque()
	fmt.Println(Queue)

}
