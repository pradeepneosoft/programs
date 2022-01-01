package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// type Data struct {
// 	state []string
// }

func main() {
	map1 := make(map[string][]string)
	map2 := make(map[string][]string)
	// var d Data
	URl := "https://raw.githubusercontent.com/bhanuc/indian-list/master/state-city.json"
	req, err := http.Get(URl)
	fmt.Println(err)
	body, err := io.ReadAll(req.Body)
	fmt.Println(err)
	err = json.Unmarshal(body, &map1)
	for i, j := range map1 {
		if string(i[0]) == "M" {
			var temp []string
			for _, q := range j {
				if string(q[0]) == "A" {
					temp = append(temp, q)
					map2[i] = temp
				}
			}
			// map2[i] = temp

		}
	}
	fmt.Println(map2)
}
