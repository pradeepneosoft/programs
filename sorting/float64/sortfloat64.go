package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []float64{0.0, 1.23, 4.56, 6.22, 1.98, 34.11, 22, 4.22}
	fmt.Println(s)

	ok := sort.Float64sAreSorted(s)
	fmt.Println(ok)
	if !ok {
		sort.Float64s(s)
	}
	fmt.Println(s)
}
