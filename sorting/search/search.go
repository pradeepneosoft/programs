package main

import (
	"fmt"
	"sort"
)

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return a[i] < a[j] }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	intSlice := []int{10, 3, 6, 10, 15, 21, 38, 26, 25, 45}
	a := 100
	pos := sort.SearchInts(intSlice, a)
	fmt.Println(pos)
	d := []int{45, 38, 26, 25, 21, 15, 10, 10, 6, 3}
	e := 25

	i := sort.Search(len(d), func(i int) bool { return d[i] <= e })
	fmt.Println(i)
	if i < len(d) && d[i] == e {
		fmt.Println(i)
	}

	elements := []string{"ruby", "python", "java", "go", "z"}

	// Sort the elements by length.
	sort.Sort(ByLen(elements))

	// Print results.
	fmt.Println(elements)
}
