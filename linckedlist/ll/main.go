package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}
type SL struct {
	Len  int
	Head *Node
}

func initList() *SL {
	return &SL{}
}
func (s *SL) AddFront(data interface{}) {
	node := &Node{
		Data: data,
	}
	if s.Head == nil {
		s.Head = node
	} else {
		node.Next = s.Head
		s.Head = node
	}
	s.Len++
	return
}
func (s *SL) AddBack(data interface{}) {
	node := &Node{
		Data: data,
	}

	if s.Head == nil {
		s.Head = node
	} else {
		current := s.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node

	}

	s.Len++
	return
}
func (s *SL) RemoveFront() error {
	if s.Head == nil {
		return errors.New("Empty List")

	}
	s.Head = s.Head.Next
	s.Len--
	return nil
}
func (s *SL) RemoveBack() error {
	if s.Head == nil {
		return errors.New("Empty List")
	}
	var pre *Node
	current := s.Head
	for current.Next != nil {
		pre = current
		current = current.Next

	}
	if pre != nil {
		pre.Next = nil
	} else {
		s.Head = nil
	}
	s.Len--
	return nil
}
func (s *SL) Size() int {
	return s.Len
}
func (s *SL) Travese() error {
	if s.Head == nil {
		return errors.New("List empty")
	}
	current := s.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
	return nil

}

func main() {
	sl := initList()
	for {

		fmt.Println("Linkedlist application")
		fmt.Println("Please Select Operation")
		fmt.Println("1  >> Add at Front")
		fmt.Println("2  >> Add at Back")

		if sl.Size() > 0 {
			fmt.Println("3  >> Remove at Front")
			fmt.Println("4  >> Remove at Back")
			fmt.Println("5  >> Traverse")

		}
		fmt.Println("0  >> Exit")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			{
				var data int

				fmt.Println("Enyter data to add front")

				fmt.Scan(&data)
				sl.AddFront(data)
			}
		case 2:
			{
				fmt.Println("Enyter data to add Back")

				var data interface{}

				fmt.Scan(&data)
				sl.AddBack(data)
			}
		case 3:
			{
				sl.RemoveFront()
			}
		case 4:
			{
				sl.RemoveBack()
			}
		case 5:
			{
				sl.Travese()
			}
		case 0:
			{
				return
			}
		default:
			fmt.Println("wrong choice")
		}
	}
}
