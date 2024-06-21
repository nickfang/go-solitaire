package game

import (
	"fmt"
)

func (g Game) Print() {
	fmt.Println("ID: ", g.Id)
	g.Stacks.Print()
	g.Board.Print()
	g.Cards.Print()
	fmt.Printf("Current Card Index: %d", g.CurrentCardIndex)
	fmt.Println()
	fmt.Printf("Flip Count: %d", g.FlipCount)
	fmt.Println()

}
