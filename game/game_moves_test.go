package game

import (
	"testing"
)

func TestNextDeckCard(t *testing.T) {
	g := NewGame("")

	// NextDeckCard behaves normally in standard cases
	err := g.NextDeckCard()
	if err != nil {
		t.Error("Expected no error.", err)
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
		t.Error("Expected no error.", err)
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
		t.Error("Expected no error.", err)
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
		t.Error("Expected no error.", err)
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
		t.Error("Expected error: invalid move.", err)
	}

	// MoveFromDeckToBoard returns an error if the move is invalid
	err = g.MoveFromDeckToBoard(3)
	if err.Error() != "invalid move" {
		t.Error("Expected error: invalid move.", err)
	}
}

func TestMoveFromDeckToStacks(t *testing.T) {
	g := NewGame("")
	g.Cards.TestingShuffle()
	g.DealBoard()

	error := g.MoveFromDeckToStacks()
	if error != nil {
		t.Error("Expected no error.", error)
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

func TestMoveFromBoardToStacks(t *testing.T) {
	g := NewGame("")
	g.Cards.TestingShuffle()
	g.DealBoard()

	for i := 0; i < 8; i++ {
		g.MoveFromDeckToStacks()
		g.MoveFromDeckToStacks()
		g.MoveFromDeckToStacks()
		g.NextDeckCard()
	}

	// MoveFromBoardToStacks behaves normally in standard cases
	err := g.MoveFromBoardToStacks(6)
	if err != nil {
		t.Error("Expected no error.", err)
	}
	if len(g.Stacks[3]) != 7 {
		t.Error("Expected seven card in stack 3.", len(g.Stacks[3]))
	}
	if g.Stacks[3][0].Value != 1 || g.Stacks[3][0].Suit != "Diamonds" {
		t.Error("Expected 5 of Clubs in stack 3.", g.Stacks[3][0])
	}
	if len(g.Board[6]) != 6 {
		t.Error("Expected six cards in column six.", len(g.Board[6]))
	}
	if g.Board[6][5].Value != 7 || g.Board[6][5].Suit != "Clubs" {
		t.Error("Expected 7 of Clubs in column 6.", g.Board[6][5])
	}

	// MoveFromBoardToStacks returns an error if the move is invalid
	err = g.MoveFromBoardToStacks(3)
	if err.Error() != "invalid move" {
		t.Error("Expected error: invalid move.", err.Error())
	}

	// MoveFromBoardToStacks returns an error if the move is invalid
	g.Board[0] = g.Board[0][:0]
	err = g.MoveFromBoardToStacks(0)
	if err.Error() != "nothing to move" {
		t.Error("Expected error: nothing to move.", err.Error())
	}
}

func TestMoveFromColumnToColumn(t *testing.T) {
	g := NewGame("")
	g.Cards.TestingShuffle()
	g.DealBoard()

	// MoveColumnToColumn behaves normally in standard cases
	err := g.MoveFromColumnToColumn(5, 4)
	if err != nil {
		t.Error("Expected no error.", err)
	}
	if len(g.Board[5]) != 5 {
		t.Error("Expected five cards in column 5.", len(g.Board[5]))
	}
	if g.Board[5][4].Value != 8 || g.Board[5][4].Suit != "Diamonds" || g.Board[5][4].Shown != true {
		t.Error("Expected 8 of Diamonds to be shown in column 5.", g.Board[5][4])
	}
	if len(g.Board[5]) != 5 {
		t.Error("Expected five cards in column 5.", len(g.Board[5]))
	}
	if g.Board[4][5].Value != 7 || g.Board[4][5].Suit != "Hearts" || g.Board[4][5].Shown != true {
		t.Error("Expected 7 of Hearts is shown in column 4.", g.Board[4][5])
	}

	// MoveColumnToColumn returns an error if the move is invalid
	err = g.MoveFromColumnToColumn(5, 5)
	if err.Error() != "invalid board move" {
		t.Error("Expected error: invalid board move.", err.Error())
	}

	// MoveColumnToColumn returns an error if the move is invalid
	g.Board[0] = g.Board[0][:0]
	err = g.MoveFromColumnToColumn(0, 1)
	if err.Error() != "invalid board move" {
		t.Error("Expected error: invalid board move.", err.Error())
	}
}
