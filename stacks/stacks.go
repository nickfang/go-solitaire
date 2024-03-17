package stacks

import (
	"fmt"
	"solitaire/deck"
)

// Stacks are where cards are pu starting with aces.
// the order of suits for the stacks are Spades, Hearts, Clubs and Diamonds.
// stacks[0] are all the Spades, stacks[1] are all the Heards ...
type Stacks [][]deck.Card

func (stacks1 Stacks) IsEqual(stacks2 Stacks) bool {
	for i := range stacks1 {
		stack1 := stacks1[i]
		stack2 := stacks2[i]
		fmt.Print(stack1)
		if !stack1[i].IsEqual(stack2[i]) {
			return false
		}
	}
	return true
}

func NewStacks() Stacks {
	stacks := Stacks{}
	for i := 0; i < 4; i++ {
		stacks = append(stacks, []deck.Card{})
	}
	return stacks
}

func (s *Stacks) MoveToStack(suitIndex int, card deck.Card) {
	if suitIndex >= 0 && suitIndex < 4 {
		(*s)[suitIndex] = append((*s)[suitIndex], card)
	} else {
		fmt.Printf("suitIndex out of range.")
		// handle out of bounds error.
	}
}

func (s Stacks) Display() {
	fmt.Print("     [")
	for _, stack := range s {
		numCards := len(stack)
		if numCards == 0 {
			fmt.Print("    ")
		} else {
			stack[numCards-1].Display()
		}
	}
	fmt.Println("]")
}
