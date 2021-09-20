package main

import "fmt"

func main() {
	myslice := make([]int, 3, 3)
	fmt.Println("Slice:", myslice[0])

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
	fmt.Println()
	myslice = append(myslice, 1)
	fmt.Println("Slice:", myslice)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
	fmt.Println()
}
