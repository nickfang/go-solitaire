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

