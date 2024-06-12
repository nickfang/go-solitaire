package game

import (
	"fmt"
)

func (g Game) DisplayCurrentCard() string {
	card, error := g.getCurrentCard()
	if error != nil {

		if error.Error() == "no cards in the deck" {
			return ""
		}
		return error.Error()
	}
	return card.DisplayMini
}

func (g Game) DisplayCards() {
	fmt.Println(g.CurrentCardIndex+1, g.DisplayCurrentCard(), len(g.Cards)-g.CurrentCardIndex)
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
	g.Stacks.Print()
	g.Board.Print()
	g.Cards.Print()
	fmt.Printf("Current Card Index: %d", g.CurrentCardIndex)
	fmt.Println()
	fmt.Printf("Flip Count: %d", g.FlipCount)
	fmt.Println()

}
