package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start go routines")
	wg.Add(3)
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	wg.Wait()
	fmt.Println("ending program")
}
func responseSize(url string) {
	defer wg.Done()

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
