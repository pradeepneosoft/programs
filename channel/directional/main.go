package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)

	receive(c)
	fmt.Println("completed")
}
func send(s chan<- int) {
	for i := 0; i < 5; i++ {
		s <- i
	}
	close(s)
}
func receive(r <-chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(<-r)
	}

}
