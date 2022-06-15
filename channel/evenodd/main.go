package main

import (
	"fmt"
)

func main() {
	// c := make(chan int)
	even := make(chan int)
	odd := make(chan int)
	q := make(chan bool)
	go send(even, odd, q)
	rec(even, odd, q)
	fmt.Println("wait")

}
func send(even chan<- int, odd chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	q <- true
	close(even)
	close(odd)
}
func rec(even <-chan int, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case o := <-even:
			fmt.Println("from even:", o)
		case e := <-odd:
			fmt.Println("from odd:", e)
		case <-quit:
			return
		}

	}
}
