package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	mutex    sync.Mutex
	counter  int32
	wg       sync.WaitGroup
	counter1 int32
)

func main() {
	wg.Add(80)
	for i := 0; i < 40; i++ {
		go show()
	}
	for i := 0; i < 40; i++ {
		go countcalls()
	}
	wg.Wait()
}
func show() {
	defer wg.Done()
	atomic.AddInt32(&counter, 1)
	fmt.Println("counter :", counter)
}
func countcalls() {
	defer wg.Done()
	mutex.Lock()
	{
		counter1++
	}
	mutex.Unlock()
	runtime.Gosched()
	fmt.Println("counter 1:", counter1)
}
