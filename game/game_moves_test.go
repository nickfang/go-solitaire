package game

import (
	"testing"
)

func TestNextDeckCard(t *testing.T) {
	g := NewGame("")

	// NextDeckCard behaves normally in standard cases
	err := g.NextDeckCard()
	if err != nil {
		t.Error("Expected no error.")
	}
	if g.CurrentCardIndex != 5 {
		t.Error("Expected current card index to be 5.", g.CurrentCardIndex)
	}
	if g.Cards[2].Shown == true {
		t.Error("Expected old card to be hidden.", g.Cards[2].Shown)
	}

	// NextDeckCard returns an error if the next card is the current card and hides and shows cards correctly
	g.Cards = g.Cards[:5]
	g.CurrentCardIndex = 2
	err = g.NextDeckCard()
	if err.Error() != "end of deck" {
		t.Error("Expected error: end of deck.", err.Error())
	}
	g.CurrentCardIndex = 1
	err = g.NextDeckCard()
	if err != nil {
		t.Error("Expected no error.")
	}
	if g.Cards[g.CurrentCardIndex].Shown != true {
		t.Error("Expected current card to be shown.", g.Cards[g.CurrentCardIndex].Shown)
	}
	if g.Cards[1].Shown == true {
		t.Error("Expected old card to be hidden.", g.Cards[1].Shown)
	}

	// Test special cases when there's 2 or less cards.
	g.Cards = g.Cards[:2]
	g.CurrentCardIndex = 0
	err = g.NextDeckCard()
	if err != nil {
		t.Error("Expected no error.")
	}
	if g.CurrentCardIndex != 1 {
		t.Error("Expected current card index to be 1 if there are only 2 cards left.", g.CurrentCardIndex)
	}
	g.Cards = g.Cards[:1]
	g.CurrentCardIndex = 0
	err = g.NextDeckCard()
	if err.Error() != "end of deck" {
		t.Error("Expected error: end of deck.", err.Error())
	}
	if g.CurrentCardIndex != 0 {
		t.Error("Expected current card index to be 0 if there are only 2 cards left.", g.CurrentCardIndex)
	}
	g.Cards = g.Cards[:0]
	err = g.NextDeckCard()
	if err.Error() != "no more cards in the deck" {
		t.Error("Expected error message not shown.")
	}

}

func TestMoveFromDeckToBoard(t *testing.T) {
	g := NewGame("")
	g.DealBoard()

	// using unshuffled deck
	err := g.MoveFromDeckToBoard(3)
	if err != nil {
		t.Error("Expected no error.")
	}
	if len(g.Board[3]) != 5 {
		t.Error("Expected five cards in column 3.", len(g.Board[3]))
	}
	if g.Board[3][4].Value != 5 || g.Board[3][4].Suit != "Clubs" {
		t.Error("Expected 5 of Clubs in column 3.", g.Board[3][4])
	}
	if g.CurrentCardIndex != 1 {
		t.Error("Expected current card index to be 1.", g.CurrentCardIndex)
	}

	// MoveFromDeckToBoard returns an error if the move is invalid
	err = g.MoveFromDeckToBoard(0)
	if err.Error() != "invalid move" {
		t.Error("Expected error: invalid move.", err.Error())
	}

	// MoveFromDeckToBoard returns an error if the move is invalid
	err = g.MoveFromDeckToBoard(3)
	if err.Error() != "invalid move" {
		t.Error("Expected error: invalid move.", err.Error())
	}
}

func TestMoveFromDeckToStacks(t *testing.T) {
	g := NewGame("")
	g.Cards.TestingShuffle()
	g.DealBoard()

	error := g.MoveFromDeckToStacks()
	if error != nil {
		t.Error("Expected no error.")
	}
	if len(g.Stacks[0]) != 1 {
		t.Error("Expected one card in stack 0.", len(g.Stacks[0]))
	}
	if g.Stacks[0][0].Value != 1 || g.Stacks[0][0].Suit != "Spades" {
		t.Error("Expected 1 of Spades in stack 0.", g.Stacks[0][0])
	}
	if g.Cards[g.CurrentCardIndex].Shown != true {
		t.Error("Expected current card in deck to be shown.", g.Cards[g.CurrentCardIndex].Shown)
	}
	if g.CurrentCardIndex != 1 {
		t.Error("Expected current card index to be 1.", g.CurrentCardIndex)
	}

}
