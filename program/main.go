package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9}
	key := 6
	count := 0
	temp := map[int]int{}
	for _, j := range a {
		if key == j {
			count++
		}
		temp[j]++

	}
	fmt.Println(count)
	fmt.Println(temp)
}
