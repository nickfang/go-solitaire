package main

import (
	"fmt"
	"solitaire/game"
	// "solitaire/main"
	"testing"
)

func TestDealBoard(t *testing.T) {
	g := game.NewGame()
	fmt.Println(g)
	g.DealBoard()
	// g.SetDebug(true)
	// g.board.Display()
}
