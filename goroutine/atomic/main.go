package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	wg      sync.WaitGroup
	counter int32
)

func main() {
	fmt.Println("start go routines")
	wg.Add(3)
	go responseSize("https")
	go responseSize("http")
	go responseSize("http2")
	wg.Wait()
	fmt.Println("ending program :", counter)
}
func responseSize(url string) {
	defer wg.Done()
	for range url {
		atomic.AddInt32(&counter, 1)
		// counter++
		runtime.Gosched()
		fmt.Println(runtime.NumCPU())

	}

}
