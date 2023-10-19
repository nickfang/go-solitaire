package main_test

import (
	"github.com/nickfang/go-solitaire/game"
	"github.com/nickfang/go-solitaire/main"
	"testing"
)

func TestDealBoard(t *testing.T) {
	g := NewGame()
	g.cards, g.board, g.currentCardIndex = g.dealBoard()
	g.setDebug(true)
	g.board.display()
}
