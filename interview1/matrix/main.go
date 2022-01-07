package main

import "fmt"

func main() {
	matrix(3)
}
func matrix(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || (i+j+1) == n {
				fmt.Print("x")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
}
