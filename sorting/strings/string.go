package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"sdfsd", "sdfsdfsdf", "aaaa", "fghfghfhgf", "opopopop"}
	fmt.Println(s)
	ok := sort.StringsAreSorted(s)
	fmt.Println(ok)

	if !ok {
		sort.Strings(s)
		fmt.Println(s)
	}
}
