package stacks_test

import (
	"solitaire/stacks"
	"testing"
)

func TestNewStacks(t *testing.T) {
	s := stacks.NewStacks()
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
