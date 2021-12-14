package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{6, 4, 3, 2, 6, 8, 4, 3, 22, 56, 7, 8, 99}
	fmt.Println(s)

	isorted := sort.IsSorted(s)
	fmt.Println(isorted)
	ok := sort.IntsAreSorted(s)
	fmt.Println(ok)

	if !ok {
		sort.Ints(s)
	}

	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
