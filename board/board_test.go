package board_test

import (
	"solitaire/board"
	"solitaire/deck"
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

func TestGetLastCard(t *testing.T) {
	b := board.NewBoard()
	index, card := b.GetLastCard(0)
	if index != 0 {
		t.Error("Index should be 0, but is ", index)
	}
	if card.Value != 0 {
		t.Error("Card should be 0, but is ", card.Value)
	}
	card1 := deck.Card{false, false, 1, "Spades", "", ""}
	b[1] = append(b[1], card1)
	index, card = b.GetLastCard(1)
	if index != 0 {
		t.Error("Index should be 0, but is ", index)
	}
	if card.Value != 1 {
		t.Error("Card should be 1, but is ", card.Value)
	}
}
