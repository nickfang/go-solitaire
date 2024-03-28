package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	g := NewGame()

	if len(g.Cards) != 52 {
		t.Error("Wrong number of cards in the deck.")
	}
}

func TestNextDeckCard(t * testing.T) {
	g := NewGame()
	g.Cards = g.Cards[:0]
	err := g.NextDeckCard()
	if err.Error() != "no more cards in the deck" {
		t.Error("Expected error message not shown.")
	}
}