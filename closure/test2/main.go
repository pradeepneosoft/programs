package main

import "fmt"

func newCounter(gc int) func() int {
	//gc := 0

	return func() int {
		fmt.Println("gc in newCounter", gc)

		gc += 1
		return gc
	}
}
func main() {

	counter := newCounter(10)
	fmt.Println(counter())
	fmt.Println(counter())

}
