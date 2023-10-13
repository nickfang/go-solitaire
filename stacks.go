package main

import "fmt"

// Stacks are where cards are pu starting with aces.
// the order of suits for the stacks are Spades, Hearts, Clubs and Diamonds.
// stacks[0] are all the Spades, stacks[1] are all the Heards ...
type stacks [][]card

func newStacks() stacks {
	stacks := stacks{}
	for i := 0; i < 4; i++ {
		stacks = append(stacks, []card{})
	}
	return stacks
}

func (s stacks) display() {
	fmt.Println(s)
}
