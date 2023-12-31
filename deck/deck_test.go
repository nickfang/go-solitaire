package deck_test

import (
	"testing"

	"solitaire/deck"
)

func TestGetCardColor(t *testing.T) {

}

func TestNewDeck(t *testing.T) {
	deck := deck.NewDeck()
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
		if c.Color != color {
			t.Error("Color was not set correctly.")
		}

		// Suits set correctly
		if index < 13 {
			if c.Suit != "Spades" {
				t.Error("Spades was not set correctly.")
			}
		}
		if index >= 13 && index < 26 {
			if c.Suit != "Hearts" {
				t.Error("Hearts was not set correctly.")
			}
		}
		if index >= 26 && index < 39 {
			if c.Suit != "Clubs" {
				t.Error("Clubs was not set correctly.")
			}
		}
		if index >= 39 && index < 52 {
			if c.Suit != "Diamonds" {
				t.Error("Diamonds was not set correctly.")
			}
		}

		// all c.shown should be initiated to false
		if c.Shown != false {
			t.Error("Show should be initialted to false.")
			break
		}
	}
}
