package main

import (
	"fmt"
	"time"
)

func main() {
	// str := "April, 01 2020 00:00:00 +0530"
	// t, err := time.Parse(time.RFC822, str)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t)

	timeT, err := time.Parse("2006-01-02", "April, 01 2020 00:00:00 +0530")
	fmt.Println(err)
	fmt.Println(timeT)
}
