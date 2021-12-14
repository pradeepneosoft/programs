package main

import "fmt"

type notifer interface {
	notify()
}
type user struct {
	name string
	age  int
}

func (u user) notify() {
	fmt.Printf("mainl to user with name  %v and age %v \n", u.name, u.age)
}
func (u admin) notify() {
	fmt.Printf("mainl to user with name  %v and age %v \n", u.name, u.age)
}

type admin struct {
	user
	department string
}

func main() {
	a := admin{
		department: "account",
		user: user{
			name: "abc",
			age:  40,
		},
	}
	a.user.notify()
	a.notify()
	sendntification(a)
	sendntification(&a.user)

}
func sendntification(n notifer) {
	n.notify()
}
