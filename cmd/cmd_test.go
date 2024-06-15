package main

import (
	"fmt"
	"solitaire/game"

	"testing"
)

func TestDealBoard(t *testing.T) {
	g := game.NewGame()
	fmt.Println(g)
	g.DealBoard()
	// g.SetDebug(true)
	// g.board.Display()
}

func TestGetCardDisplay(t *testing.T) {
	cardDisplays := []string{}
	expected := []string{" A♠", " 2♠", " 3♠", " 4♠", " 5♠", " 6♠", " 7♠", " 8♠", " 9♠", "10♠", " J♠", " Q♠", " K♠", " A♥", " 2♥", " 3♥", " 4♥", " 5♥", " 6♥", " 7♥", " 8♥", " 9♥", "10♥", " J♥", " Q♥", " K♥", " A♣", " 2♣", " 3♣", " 4♣", " 5♣", " 6♣", " 7♣", " 8♣", " 9♣", "10♣", " J♣", " Q♣", " K♣", " A♦", " 2♦", " 3♦", " 4♦", " 5♦", " 6♦", " 7♦", " 8♦", " 9♦", "10♦", " J♦", " Q♦", " K♦"}

	for _, suit := range []string{"Spades", "Hearts", "Clubs", "Diamonds"} {
		for j := 1; j <= 13; j++ {
			cardDisplays = append(cardDisplays, getCardDisplay(j, suit))
		}
	}
	for i, cardDisplay := range cardDisplays {
		if cardDisplay != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], cardDisplay)
		}
	}
}
