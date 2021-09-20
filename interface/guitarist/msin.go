package main

import "fmt"

type Guitarist interface {
	PlayGuitar()
}
type BaseGuitarist struct {
	Name string
}
type Accusticguitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays base guitar ", b.Name)
	fmt.Println()

}
func (a Accusticguitarist) PlayGuitar() {
	fmt.Printf("%s plays accoustic guitar", a.Name)
	fmt.Println()

}
func main() {
	var guitarists []Guitarist

	var player BaseGuitarist
	player.Name = "sarkar"
	player.PlayGuitar()
	var player2 Accusticguitarist
	player2.Name = "Raaj"
	player2.PlayGuitar()
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)
	for _, j := range guitarists {

		j.PlayGuitar()
	}
}
