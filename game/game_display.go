package game

import (
	"fmt"
)

func (g Game) DisplayCurrentCard() string {
	return g.getCurrentCard().DisplayMini
}

func (g Game) DisplayCards() {
	fmt.Println(g.CurrentCardIndex, g.DisplayCurrentCard(), len(g.Cards)-g.CurrentCardIndex)
}

func (g Game) Display() {
	g.Stacks.Display()
	g.Board.Display()
	g.DisplayCards()
	if g.Debug {
		fmt.Println(g.GetDeckMoves())
	}
}

func (g Game) DisplayHints() {
	deckMoves := g.GetDeckHints()
	stackMoves := g.GetStackHints()
	fmt.Println("Moves:", append(deckMoves, stackMoves...))
}

func (g Game) Print() {
	fmt.Println("Stacks:")
	g.Stacks.Display()
	fmt.Println("Board:")
	g.Board.Print()
	fmt.Println("Deck:")
	g.Cards.Print()
	fmt.Printf("Current Card Index: %d", g.CurrentCardIndex)
	fmt.Println()
}
