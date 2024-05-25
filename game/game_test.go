package game

import (
	"solitaire/testutils"
	"testing"
)

func TestIsEqual(t *testing.T) {
	game1 := NewGame()
	game2 := NewGame()
	if !game1.IsEqual(game2) {
		t.Error("Initial game state should be equal.")
	}

	game1.Cards.RandomShuffle()
	game2.Cards.PerfectShuffle()
	if game1.IsEqual(game2) {
		t.Error("Shuffled decks should not be equal.")
	}

	for i := 0; i < NumCards; i++ {
		game1.Cards.RemoveCard(0)
	}
	if game1.IsEqual(game2) {
		t.Error("Empty deck should not be equal to another deck.")
	}

	game1 = NewGame()
	for i := 0; i < NumCards; i++ {
		game1.MoveFromDeckToStacks()
	}
	if game1.IsEqual(game2) {
		t.Error("Full Stack should not be equal to an empty stack.")
	}
}

func TestNewGame(t *testing.T) {
	g := NewGame()

	if len(g.Cards) != 52 {
		t.Error("Wrong number of cards in the deck.")
	}
}

func TestNextDeckCard(t *testing.T) {
	g := NewGame()

	g.Cards = g.Cards[:0]
	err := g.NextDeckCard()
	if err.Error() != "no more cards in the deck" {
		t.Error("Expected error message not shown.")
	}

}

func TextCheckMove(t *testing.T) {

}

func TestDeepCopy(t *testing.T) {
	game := NewGame()
	gameShallow := game
	gameDeep := game.DeepCopy()

	// Checking each part of the game Cards, Board, Stacks and CurrentCardIndex
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
}
