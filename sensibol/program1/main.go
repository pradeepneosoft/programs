package main

import (
	"fmt"
	"strings"
)

func main() {

	map1 := make(map[string]string)

	map1["name"] = "pradeep kachraram bora"
	map1["age"] = "25"
	map1["nationality"] = "indian"

	fmt.Println(map1)

	// for i, j := range map1 {name
	// if i == "name" {
	tslice := strings.Split(map1["name"], " ")
	tempstring := ""
	for _, q := range tslice {
		tempstring += strings.ToUpper(string(q[0])) + q[1:] + " "
		// fmt.Println(tempstring)

	}
	map1["name"] = tempstring
	// map1["name1"]=tslice[]
	fmt.Println(map1)
	// }
	// }
}
