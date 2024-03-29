package stacks

import (
	"fmt"
	"solitaire/deck"
)

// Stacks are where cards are put starting with aces.
// the order of suits for the stacks are Spades, Hearts, Clubs and Diamonds.
// stacks[0] are all the Spades, stacks[1] are all the Heards ...
type Stack deck.Cards
type Stacks []deck.Cards

func (stacks1 Stacks) IsEqual(stacks2 Stacks) bool {
	for i := 0; i < 4; i++ {
		stack1 := stacks1[i]
		stack2 := stacks2[i]
		if len(stack1) != len(stack2) {
			return false
		}
		// if both length are 0, don't need to check contents
		if len(stack1) == 0 && len(stack2) == 0 {
			continue
		}
		if (!stack1.IsEqual(stack2)) {
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

func (s *Stacks) MoveToStack(card deck.Card) {
	suitIndex := -1
	for index, suit := range deck.CardSuits {
		if suit == card.Suit {
			suitIndex = index
			break
		}
	}
	if suitIndex == -1 {
		fmt.Printf("Invalid suit in card.  This card was created incorrectly.  This should never happen.")
		return
	}
	if len((*s)[suitIndex]) == 0 {
		if card.Value != 1 {
			fmt.Printf("Only an ace can be moved to an empty stack.")
			return
		}
	} else if ((*s)[suitIndex][len((*s)[suitIndex])-1].Value + 1 != card.Value) {
		fmt.Printf("No valid spots in the stack for this card.")
		return
	}
	(*s)[suitIndex] = append((*s)[suitIndex], card)
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
