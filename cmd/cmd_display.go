package main

import (
	"fmt"
	"solitaire/game"
	"solitaire/game/board"
	"solitaire/game/deck"
	"solitaire/game/stacks"
)

// CLI Display
func DisplayBoard(b board.Board) {
	displayBoard := [7][19]deck.Card{}
	maxLen := 0 // add a space so the board isn't cramped with the deck.
	for i, column := range b {
		if len(column) > maxLen {
			maxLen = len(column)
		}
		for j, card := range column {
			displayBoard[i][j] = card
		}
	}
	for y := 0; y < maxLen; y++ {
		for x := 0; x < 7; x++ {
			displayBoard[x][y].Display()
		}
		fmt.Println()
	}
}

// CLI display
func DisplayCard(c deck.Card) {
	if c.Value == 0 {
		fmt.Print("    ")
		return
	}
	if c.Shown || c.Debug {
		fmt.Print(c.DisplayMini)
	} else {
		fmt.Print("  * ")
	}
}

// CLI Display
func DisplayStacks(s stacks.Stacks) {
	fmt.Print("     [")
	for _, stack := range s {
		numCards := len(stack)
		if numCards == 0 {
			fmt.Print("    ")
		} else {
			stack[numCards-1].Display()
		}
	}
	fmt.Println("]")
}

func DisplayCurrentCard(g game.Game) string {
	card, error := g.GetCurrentCard()
	if error != nil {

		if error.Error() == "no cards in the deck" {
			return ""
		}
		return error.Error()
	}
	return card.DisplayMini
}

func DisplayCards(g game.Game) {
	fmt.Println(g.CurrentCardIndex+1, g.DisplayCurrentCard(), len(g.Cards)-g.CurrentCardIndex)
}

func DisplayGame(g game.Game) {
	g.Stacks.Display()
	g.Board.Display()
	g.DisplayCards()
	if g.Debug {
		fmt.Println(g.GetDeckMoves())
	}
}

func DisplayHints(g game.Game) {
	hints := g.GetDeckHints()
	hints = append(hints, g.GetStackHints()...)
	hints = append(hints, g.GetBoardHints()...)
	fmt.Println("Moves:", hints)
}
