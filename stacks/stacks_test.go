package stacks

import (
	"solitaire/deck"
	"testing"
)


func TestNewStacks(t *testing.T) {
	s := NewStacks()
	if len(s) != 4 {
		t.Error("Stacks has the wrong number of columns.")
	}
	for i, column := range s {
		if len(column) != 0 {
			t.Error("Stacks has the wrong number of rows.")
		}
		for j, card := range column {
			if card.Value != 0 {
				t.Error("Card ", i, ", ", j, "was not initialized to 0.")
			}
		}
	}
}

func TestMoveToStack(t *testing.T) {
	s := NewStacks()

	c1, err1 := deck.NewCard(1, "Hearts", true)
	if err1 != nil {
		t.Error("error creating c1 card.")
	}
	s.MoveToStack(c1)

	c2, err2 := deck.NewCard(2, "Spades", true)
	if err2 != nil {
		t.Error("error creating c1 card.")
	}
	s.MoveToStack(c2)

}

func TestIsEqual(t *testing.T) {
	// deck := deck.NewDeck()
	stack1 := NewStacks()
	stack2 := NewStacks()
	if !stack1.IsEqual(stack2) {
		t.Error("The stacks should be equal.")
	}

}