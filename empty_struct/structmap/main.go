package main

import "fmt"

func main() {

	searchinslice(3, 4, 5, 76, 7, 3, 2, 4, 5, 6, 75, 4, 3, 21, 2, 3, 56, 5, 78)
	searchinslice1(3, 4, 5, 76, 7, 3, 2, 4, 5, 6, 75, 4, 3, 21, 2, 3, 56, 5, 78)

}
func searchinslice(d ...int) int {
	counter := 0
	mymap := map[int]string{
		1: "dsdsds",
		2: "dsdsds",
		3: "dsdsds",
	}
	for j, _ := range mymap {
		for _, q := range d {
			if j == q {
				counter++
			}
		}
	}
	fmt.Println(counter)

	return counter
}
func searchinslice1(d ...int) int {
	counter := 0
	mymap := map[int]string{
		1: "dsdsds",
		2: "dsdsds",
		3: "dsdsds",
	}
	elementMap := make(map[int]struct{})
	type demo struct{}
	for _, data := range d {
		elementMap[data] = struct{}{}
	}
	for j, _ := range mymap {
		_, ok := elementMap[j]
		if ok {
			counter++
		}
	}
	fmt.Println(counter)
	return counter
}
