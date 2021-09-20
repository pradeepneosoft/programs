package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch := make(chan string, 3)
	go write(ch)
	time.Sleep(2 * time.Second)
	for b := range ch {
		fmt.Println("rading : " + b)
	}

}

func write(ch chan string) {
	for i := 0; i < 4; i++ {
		ch <- "Hello" + strconv.Itoa(i)
		fmt.Println("witung : " + strconv.Itoa(i))
	}
	close(ch)
}
