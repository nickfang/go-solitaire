package stacks

import (
	"fmt"
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
	c := deck.NewCard(1, "Hearts", true)
	fmt.Print(c)
	s.MoveToStack(0, c)
	s.Display()
	t.Error("test")
}

func TestIsEqual(t *testing.T) {
	// deck := deck.NewDeck()
	stack1 := NewStacks()
	stack2 := NewStacks()
	if !stack1.IsEqual(stack2) {
		t.Error("The stacks should be equal.")
	}

}