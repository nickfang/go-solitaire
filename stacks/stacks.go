package stacks

import (
	"fmt"
	"solitaire/deck"
)

// Stacks are where cards are pu starting with aces.
// the order of suits for the stacks are Spades, Hearts, Clubs and Diamonds.
// stacks[0] are all the Spades, stacks[1] are all the Heards ...
type Stacks [][]deck.Card

func NewStacks() Stacks {
	stacks := Stacks{}
	for i := 0; i < 4; i++ {
		stacks = append(stacks, []deck.Card{})
	}
	return stacks
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
