package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Data struct {
	No   int
	Name int
}

func main() {
	c1 := make(chan Data)
	c2 := make(chan Data)

	go populate(c1)
	go fanout(c1, c2)
	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

func populate(c1 chan<- Data) {
	for i := 0; i < 100; i++ {
		c1 <- Data{No: i}
	}
	close(c1)
}

func fanout(c1, c2 chan Data) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 Data) {
			c2 <- dosomething(v2)
			wg.Done()

		}(v)
	}
	wg.Wait()
	close(c2)
}
func dosomething(n Data) Data {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return Data{Name: n.No + rand.Intn(1000), No: n.No}
}
