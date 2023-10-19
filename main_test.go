package main_test

import (
	"fmt"
	"solitaire/game"
	// "solitaire/main"
	"testing"
)

func TestDealBoard(t *testing.T) {
	g := game.NewGame()
	fmt.Println(g)
	// g.cards, g.board, g.currentCardIndex = g.DealBoard()
	// g.SetDebug(true)
	// g.board.Display()
}
