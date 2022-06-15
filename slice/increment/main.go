package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	n := 4
	wg.Add(4)
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println("hello world")
			defer wg.Done()

		}()
	}
	wg.Wait()
}
