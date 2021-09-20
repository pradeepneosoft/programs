package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("step 1:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("step 2: ", url)
	defer response.Body.Close()

	fmt.Println("step 3: ", url)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))

}

func main() {
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(5 * time.Second)

}
