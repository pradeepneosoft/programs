package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c <- "something"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c2 <- "Nothing"
			time.Sleep(2 * time.Second)
		}
	}()
	for {
		select {
		case r := <-c:
			fmt.Println(r)
		case r2 := <-c2:
			fmt.Println(r2)
		}

	}
}
