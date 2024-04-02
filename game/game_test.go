package game

import (
	"fmt"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := NewGame()

	if len(g.Cards) != 52 {
		t.Error("Wrong number of cards in the deck.")
	}
}

func TestNextDeckCard(t * testing.T) {
	g := NewGame()

	g.Cards = g.Cards[:0]
	err := g.NextDeckCard()
	if err.Error() != "no more cards in the deck" {
		t.Error("Expected error message not shown.")
	}

}

func TestDeepCopy(t *testing.T) {
	game := NewGame()
	gameShallow := game
	gameDeep := game.DeepCopy()

	// Checking each part of the game Cards, Board, Stacks and CurrentCardIndex
	game.MoveFromDeckToStacks()
	if !game.IsEqual(gameShallow) {
		fmt.Printf("Shallow copy should be updated when original is updated")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}
	game.Cards.RandomShuffle()
	if !game.IsEqual(gameShallow) {
		fmt.Printf("Shallow copy should be updated when original is updated")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}
	game.DealBoard()
	if !game.IsEqual(gameShallow) {
		fmt.Printf("Shallow copy should be updated when original is updated")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}
	game.NextDeckCard()
	if !game.IsEqual(gameShallow) {
		fmt.Printf("Shallow copy should be updated when original is updated")
	}
	if game.IsEqual(gameDeep) {
		t.Error("Deep copy should not be updated when original is updated.")
	}
}