package stacks

import (
	"fmt"
	"solitaire/game/deck"
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
	card1, _ := deck.NewCard(1, "Hearts", true)
	card2, _ := deck.NewCard(2, "Hearts", true)
	card3, _ := deck.NewCard(1, "Spades", true)
	card4, _ := deck.NewCard(1, "Diamonds", true)
	card5, _ := deck.NewCard(2, "Diamonds", true)
	card6, _ := deck.NewCard(3, "Diamonds", true)
	stack1.MoveToStack(card1)
	stack1.MoveToStack(card2)
	stack1.MoveToStack(card3)
	stack1.MoveToStack(card4)
	stack1.MoveToStack(card5)
	stack1.MoveToStack(card6)

	equals1 := stack1.IsEqual(stack2)
	if equals1 {
		t.Error("stacks should not match")
	}
	stack2[1] = append(stack2[1], card1)
	equals2 := stack1.IsEqual(stack2)
	if equals2 {
		t.Error("stacks should not match")
	}
	stack2[1] = append(stack2[1], card2)
	equals3 := stack1.IsEqual(stack2)
	if equals3 {
		t.Error("stacks should not match")
	}
	stack2[0] = append(stack2[0], card3)
	equals4 := stack1.IsEqual(stack2)
	if equals4 {
		t.Error("stacks should not match")
	}
	stack2[3] = append(stack2[3], card4)
	equals5 := stack1.IsEqual(stack2)
	if equals5 {
		t.Error("stacks should not match")
	}
	stack2[3] = append(stack2[3], card5)
	equals6 := stack1.IsEqual(stack2)
	if equals6 {
		t.Error("stacks should not match")
	}
	stack2[3] = append(stack2[3], card6)
	equals7 := stack1.IsEqual(stack2)
	fmt.Println(stack1)
	fmt.Println(stack2)
	stack1.Display()
	stack2.Display()
	if !equals7 {
		t.Error("stacks should match")
	}
}

func TestGetTopCard(t *testing.T) {
	stack := NewStacks()
	card1, _ := deck.NewCard(1, "Hearts", true)
	card2, _ := deck.NewCard(1, "Spades", true)
	card3, _ := deck.NewCard(1, "Diamonds", true)
	card4, _ := deck.NewCard(2, "Diamonds", true)
	card5, _ := deck.NewCard(3, "Diamonds", true)

	stack.MoveToStack(card1)
	stack.MoveToStack(card2)
	stack.MoveToStack(card3)
	stack.MoveToStack(card4)
	stack.MoveToStack(card5)
	fmt.Println(stack)
	stacks := stack.GetTopCards()
	if stacks[0][0] != card2 {
		t.Error("top spades should be an ace")
	}
	if stacks[1][0] != card1 {
		t.Error("top hearts should be an ace")
	}
	if len(stacks[2]) != 0 {
		t.Error("there should be no top club card")
	}
	if stacks[3][0] != card5 {
		t.Error("top diamonds should be an ace")
	}
}