package stacks

import (
	"testing"
)

func (stacks1 Stacks) IsEqual(stacks2 Stacks) bool {
	for i := range stacks1 {
		stack1 := stacks1[i]
		stack2 := stacks2[i]
		if !stack1[i].insEqual(stack2[i]) {
			return false
		}
	}
	return true
}

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
