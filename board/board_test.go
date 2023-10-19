package board_test

import (
	"solitaire/board"
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := board.NewBoard()
	if len(b) != 7 {
		t.Error("Board has the wrong number of columns.")
	}
	for i, column := range b {
		if len(column) != 0 {
			t.Error("Board has the wrong number of rows.")
		}
		for j, card := range column {
			if card.Value != 0 {
				t.Error("Card ", i, ", ", j, "was not initialized to 0.")
			}
		}
	}
}
