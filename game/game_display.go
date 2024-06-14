package game

import (
	"fmt"
)

func (g Game) Print() {
	g.Stacks.Print()
	g.Board.Print()
	g.Cards.Print()
	fmt.Printf("Current Card Index: %d", g.CurrentCardIndex)
	fmt.Println()
	fmt.Printf("Flip Count: %d", g.FlipCount)
	fmt.Println()

}
