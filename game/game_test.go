package game

import (
	"solitaire/game/deck"
	"testing"
)

func TestCheckMove(t *testing.T) {
	g := NewGame("")
	card1, _ := deck.NewCard(1, "Spades", false)
	card2, _ := deck.NewCard(2, "Spades", false)
	card3, _ := deck.NewCard(2, "Hearts", false)
	card4, _ := deck.NewCard(13, "Spades", false)
	card5, _ := deck.NewCard(13, "Hearts", false)

	g.Board[0] = append(g.Board[0], card1)
	g.Board[1] = append(g.Board[1], card2)
	g.Board[2] = append(g.Board[2], card3)
	g.Board[3] = append(g.Board[3], card4)
	g.Board[4] = append(g.Board[4], card5)
	g.Board[5] = append(g.Board[5], deck.Card{})
	if checkMove(g.Board[0][0], g.Board[1][0]) {
		t.Error("Move should be invalid.")
	}
	if !checkMove(g.Board[0][0], g.Board[2][0]) {
		t.Error("Move should be valid.")
	}
	if !checkMove(g.Board[3][0], g.Board[5][0]) {
		t.Error("Move should be valid.")
	}
	if !checkMove(g.Board[4][0], g.Board[5][0]) {
		t.Error("Move should be valid.")
	}
}

func TestNewGame(t *testing.T) {
	g := NewGame("")

	if g.Id != "main" {
		t.Error("Expected game id to be 'main'.")
	}
	if len(g.Cards) != 52 {
		t.Error("Wrong number of cards in the deck.")
	}
	if g.CurrentCardIndex != DefaultFlipCount-1 {
		t.Error("Expected current card index to be 2.", g.CurrentCardIndex)
	}
	if g.FlipCount != DefaultFlipCount {
		t.Error("Expected flip count to be 3.", g.FlipCount)
	}

	g = NewGame("test")
	if g.Id != "test" {
		t.Error("Expected game id to be 'test'.")
	}

}

func TestGetCurrentCard(t *testing.T) {
	game := NewGame("")
	card, err := game.GetCurrentCard()
	if err != nil {
		t.Error("Expected no error.", err)
	}
	if card != game.Cards[2] {
		t.Error("Expected current card to be the third card in the deck.")
	}

	game.CurrentCardIndex = -1
	card, err = game.GetCurrentCard()
	if err != nil {
		t.Error("Expected no error.", err)
	}
	if card.Value != 0 || card.Suit != "" {
		t.Error("Expected current card to be empty.")
	}

	game.CurrentCardIndex = 52
	_, err = game.GetCurrentCard()
	if err.Error() != "current card index out of range" {
		t.Error("Expected error message not shown.")
	}

	game.Cards = nil
	_, err = game.GetCurrentCard()
	if err.Error() != "deck not initialized" {
		t.Error("Expected error message not shown.")
	}

	game.Cards = deck.Cards{}
	_, err = game.GetCurrentCard()
	if err.Error() != "no cards in the deck" {
		t.Error("Expected error message not shown.")
	}
}

func TestSetFlipcount(t *testing.T) {
	g := NewGame("")

	err := g.SetFlipCount(1)
	if err != nil {
		t.Error("Expected no error.", err)
	}
	if g.FlipCount != 1 {
		t.Error("Expected flip count to be 1.", g.FlipCount)
	}

	err = g.SetFlipCount(3)
	if err != nil {
		t.Error("Expected no error.", err)
	}
	if g.FlipCount != 3 {
		t.Error("Expected flip count to be 3.", g.FlipCount)
	}

	err = g.SetFlipCount(0)
	if err.Error() != "flip count must be 1 or 3" {
		t.Error("Expected error message not shown.")
	}

	err = g.SetFlipCount(-1)
	if err.Error() != "flip count must be 1 or 3" {
		t.Error("Expected error message not shown.")
	}
}

// TODO: test doing an undo and then trying to go to the next card

// TODO: test
