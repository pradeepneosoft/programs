package main

import "fmt"

func main() {
	myslice := make([]string, 3, 3)
	fmt.Println("Slice:", myslice[2])

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
	fmt.Println()
	// Creating an array
	arr := [7]string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}
	fmt.Printf("Length of the arr: %d", len(arr))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the arr: %d", cap(arr))
	fmt.Println()

	// Display array
	fmt.Println("Array:", arr)

	// Creating a slice
	myslice = arr[1:6]

	// Display slice
	fmt.Println("Slice:", myslice)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
	fmt.Println()
	myslice = append(myslice[5:6], myslice[1:4]...)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
	fmt.Println()
	fmt.Println("Array:", arr)

	// Display slice
	fmt.Println("Slice:", myslice)
	// Display length of the slice
	fmt.Printf("Length of the arr: %d", len(arr))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the arr: %d", cap(arr))
	fmt.Println()

}
