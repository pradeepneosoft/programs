package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg      sync.WaitGroup
	counter int32
	mutex   sync.Mutex
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
		mutex.Lock()
		{
			fmt.Println(url)
			counter++
		}
		mutex.Unlock()

		runtime.Gosched()
		fmt.Println(runtime.NumGoroutine())

	}

}
