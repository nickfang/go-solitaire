package main

import (
	"testing"
)

func TestGetCardColor(t *testing.T) {

}

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Error("Deck is the wrong size.")
	}
	for index, c := range deck {
		// first 13 are black next 13 red, next 13 black and the last 13 are red
		var color string
		if index%26 < 13 {
			color = "Black"
		} else {
			color = "Red"
		}
		if c.color != color {
			t.Error("Color was not set correctly.")
		}

		// Suits set correctly
		if index < 13 {
			if c.suit != "Spades" {
				t.Error("Spades was not set correctly.")
			}
		}
		if index >= 13 && index < 26 {
			if c.suit != "Hearts" {
				t.Error("Hearts was not set correctly.")
			}
		}
		if index >= 26 && index < 39 {
			if c.suit != "Clubs" {
				t.Error("Clubs was not set correctly.")
			}
		}
		if index >= 39 && index < 52 {
			if c.suit != "Diamonds" {
				t.Error("Diamonds was not set correctly.")
			}
		}

		// all c.shown should be initiated to false
		if c.shown != false {
			t.Error("Show should be initialted to false.")
			break
		}
	}
}
