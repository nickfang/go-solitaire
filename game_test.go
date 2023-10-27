package main

import (
	// "fmt"
	// "os"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := newGame()

	if len(g.cards) != 52 {
		t.Error("Wrong number of cards in the deck.")
	}
	// fmt.Println(g.cards)
	// fmt.Println(g.board)
	// fmt.Println(g.currentCardIndex)

}
