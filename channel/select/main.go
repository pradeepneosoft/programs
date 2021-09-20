package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 3)

	ch2 := make(chan string, 3)

	go write1(ch1)
	go write2(ch2)

	select {
	case r1 := <-ch1:
		fmt.Println("from ch1 " + r1)

	case r2 := <-ch2:
		fmt.Println("from ch2 " + r2)
	}

}

func write1(ch1 chan string) {
	// for i := 0; i < 4; i++ {
	time.Sleep(3 * time.Second)

	ch1 <- "Hello"
	fmt.Println("writing 1 ")
	// }
	// close(ch1)
}

func write2(ch2 chan string) {
	// for i := 0; i < 4; i++ {

	time.Sleep(6 * time.Second)
	ch2 <- "Hello"
	fmt.Println("writing 2  ")
	// }
	// close(ch2)
}
