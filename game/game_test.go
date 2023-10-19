package game_test

import (
	"solitaire/game"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := game.NewGame()

	if len(g.Cards) != 52 {
		t.Error("Wrong number of cards in the deck.")
	}

}
