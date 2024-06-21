package board

import (
	"fmt"
	"solitaire/game/deck"
	"testing"
)

func TestIsEqual(t *testing.T) {
	board := NewBoard()
	board2 := NewBoard()
	equal := board.IsEqual(board2)
	if !equal {
		t.Error("Empty boards should be equal.")
	}

}

func TestNewBoard(t *testing.T) {
	b := NewBoard()
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
	b := NewBoard()
	index, card := b.GetLastCard(0)
	fmt.Println(index, card)
	if index != -1 {
		t.Error("Index should be -1, but is ", index)
	}
	if card.Value != 0 {
		t.Error("Card should be 0, but is ", card.Value)
	}
	t.Fail()
	card1, _ := deck.NewCard(1, "Spades", false)
	b[1] = append(b[1], card1)
	index, card = b.GetLastCard(1)
	if index != 0 {
		t.Error("Index should be 0, but is ", index)
	}
	if card.Value != 1 {
		t.Error("Card should be 1, but is ", card.Value)
	}

}
