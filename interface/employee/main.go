package main

import "fmt"

type Employee interface {
	Age() int

	Language() string
	Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " Program in go"
}
func (e Engineer) Age() int {
	return 10
}
func (e Engineer) Random() (int, error) {
	return 29, nil
}
func main() {
	// var programers []Employee
	elliot := Engineer{Name: "elliot"}
	//programers = append(programers, elliot)
	fmt.Println(elliot.Age())
	fmt.Println(elliot.Language())
	fmt.Println(elliot.Random())
	// for _, j := range programers {
	// 	j.Age()
	// 	j.Language()
	// 	j.Random()
	// }

}
