package main

import "fmt"

func main() {
	myslice := []int{1, 2, 34, 4, 5, 6, 7, 8, 9, 0, 3}
	t := evenSum(sum, myslice...)
	fmt.Println(t)
}
func sum(x ...int) int {
	n := 0
	for _, j := range x {
		n += j
	}
	return n
}
func evenSum(f func(x ...int) int, val ...int) int {
	var xi []int
	for _, v := range val {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}
