package game

import (
	"fmt"
	"solitaire/game/deck"
	"solitaire/testutils"
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

	g.Print()
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
		t.Error("Expected no error.")
	}
	if card != game.Cards[2] {
		t.Error("Expected current card to be the third card in the deck.")
	}

	game.CurrentCardIndex = -1
	card, err = game.GetCurrentCard()
	if err != nil {
		t.Error("Expected no error.")
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

func TestIsEqual(t *testing.T) {
	game1 := NewGame("")
	game2 := NewGame("")
	if !game1.IsEqual(game2) {
		t.Error("Initial game state should be equal.")
	}

	game1.Cards.RandomShuffle()
	game2.Cards.PerfectShuffle()
	if game1.IsEqual(game2) {
		t.Error("Shuffled decks should not be equal.")
	}

	for i := 0; i < deck.DeckSize; i++ {
		game1.Cards.RemoveCard(0)
	}
	if game1.IsEqual(game2) {
		t.Error("Empty deck should not be equal to another deck.")
	}

	game1 = NewGame("")
	for i := 0; i < deck.DeckSize; i++ {
		game1.MoveFromDeckToStacks()
	}
	if game1.IsEqual(game2) {
		t.Error("Full Stack should not be equal to an empty stack.")
	}
}

func TestNextDeckCard(t *testing.T) {
	g := NewGame("")

	err := g.NextDeckCard()
	if err != nil {
		t.Error("Expected no error.")
	}
	if g.CurrentCardIndex != 5 {
		t.Error("Expected current card index to be 5.", g.CurrentCardIndex)
	}

	g.Cards = g.Cards[:2]
	g.CurrentCardIndex = 0
	fmt.Println(g.Cards)
	err = g.NextDeckCard()
	if err != nil {
		t.Error("Expected no error.")
	}
	if g.CurrentCardIndex != 1 {
		t.Error("Expected current card index to be 1 if there are only 2 cards left.", g.CurrentCardIndex)
	}

	g.Cards = g.Cards[:1]
	g.CurrentCardIndex = 0
	fmt.Println(g.Cards)
	err = g.NextDeckCard()
	if err != nil {
		t.Error("Expected no error.")
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

func TestSetFlipcount(t *testing.T) {
	g := NewGame("")

	err := g.SetFlipCount(1)
	if err != nil {
		t.Error("Expected no error.")
	}
	if g.FlipCount != 1 {
		t.Error("Expected flip count to be 1.", g.FlipCount)
	}

	err = g.SetFlipCount(3)
	if err != nil {
		t.Error("Expected no error.")
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

func TextCheckMove(t *testing.T) {

}

func TestDeepCopy(t *testing.T) {
	game := NewGame("")
	gameShallow := game
	gameDeep := game.DeepCopy()
	// Checking each part of the game Cards, Board, Stacks, CurrentCardIndex and FlipCount
	game.CurrentCardIndex = 0
	testutils.AssertNoError(t, game.MoveFromDeckToStacks(), "MoveFromDeckToStacks")
	if !game.Cards[0].IsEqual(gameShallow.Cards[0]) {
		t.Error("Shallow copies behave unexpectedly.  If you need to make a copy of game, use deep copy.")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}

	testutils.AssertNoError(t, game.Cards.RandomShuffle(), "RandomShuffle")
	if game.Cards.IsEqual(gameShallow.Cards) {
		t.Error("Shallow copies behave unexpectedly.  If you need to make a copy of game, use deep copy.")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}

	game.DealBoard()
	if game.IsEqual(gameShallow) {
		t.Error("Shallow copies behave unexpectedly.  If you need to make a copy of game, use deep copy.")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}

	testutils.AssertNoError(t, game.NextDeckCard(), "NextDeckCard")
	if game.IsEqual(gameShallow) {
		t.Error("Shallow copies behave unexpectedly.  If you need to make a copy of game, use deep copy.")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}

	testutils.AssertNoError(t, game.SetFlipCount(1), "FlipCount")
	if game.IsEqual(gameShallow) {
		t.Error("Shallow copies behave unexpectedly.  If you need to make a copy of game, use deep copy.")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}
	// TODO: Test CurrentCardIndex and FlipCount
}

// TODO: test doing an undo and then trying to go to the next card

// TODO: test

func TestMoveColumnToColumn(t *testing.T) {
	g := NewGame("")
	g.Cards.TestingShuffle()
	g.DealBoard()
	// moves := []string{"54", "53"}
}
