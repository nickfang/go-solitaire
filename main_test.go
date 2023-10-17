package main

import "testing"

func TestDealBoard(t *testing.T) {
	g := newGame()
	g.cards, g.board, g.currentCardIndex = g.dealBoard()
	g.setDebug(true)
	g.board.display()
}
