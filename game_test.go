package main

import (
	"fmt"
	// "os"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := newGame()
	fmt.Println(g.cards)
	fmt.Println(g.board)
	fmt.Println(g.currentCardIndex)
	t.Log("test run")
}
