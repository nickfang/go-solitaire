package stacks

import (
	"errors"
	"fmt"
	"solitaire/game/deck"
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
		if !stack1.IsEqual(stack2) {
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

func (s *Stacks) MoveToStack(card deck.Card) error {
	suitIndex := -1
	for index, suit := range deck.CardSuits {
		if suit == card.Suit {
			suitIndex = index
			break
		}
	}
	if suitIndex == -1 {
		return errors.New("invalid suit in card - use NewCard() when creating card")
	}
	if len((*s)[suitIndex]) == 0 {
		if card.Value != 1 {
			return errors.New("only an ace can be moved to an empty stack")
		}
	} else if (*s)[suitIndex][len((*s)[suitIndex])-1].Value+1 != card.Value {
		return errors.New("invalid stack move")
	}
	(*s)[suitIndex] = append((*s)[suitIndex], card)
	return nil
}

func (s Stacks) GetTopCards() Stacks {
	// handle error if index is > 3
	// maybe find a better way to figure out which stack they are trying to access
	stackTops := NewStacks()
	for i, stack := range s {
		if len(stack) == 0 {
			continue
		}
		index := len(stack) - 1
		stackTops[i] = append(stackTops[i], stack[index])
	}
	return stackTops
}

// for debugging
func (s Stacks) Print() {
	for i, stack := range s {
		fmt.Printf("Stack %d: ", i)
		for _, card := range stack {
			card.Print()
		}
		fmt.Println()
	}
}
