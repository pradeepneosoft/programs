package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{6, 4, 3, 2, 6, 8, 4, 3, 22, 56, 7, 8, 99}
	fmt.Println(s)
	tempmap := make(map[int]struct{})
	tempslice := []int{}
	for _, j := range s {

		if _, ok := tempmap[j]; !ok {
			tempmap[j] = struct{}{}
			tempslice = append(tempslice, j)
		}
	}
	fmt.Println(tempslice)
	ok := sort.IntsAreSorted(s)
	fmt.Println(ok)

	if !ok {
		sort.Ints(s)
	}

	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	// for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
	// 	fmt.Println(s[a], "========", s[b])
	// 	s[a], s[b] = s[b], s[a]

	// }
	fmt.Println(s)
}
