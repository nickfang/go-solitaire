package deck

import (
	"fmt"
	"testing"
)

func TestIsEqual(t *testing.T) {
	card1 := Card{
		false,
		false,
		1,
		"Spades",
		getCardColor("Spades"),
		"",
	}
	card2 := Card{
		false,
		false,
		1,
		"Spades",
		getCardColor("Spades"),
		"",
	}
	if !card1.IsEqual(card2) {
		t.Error("IsEqual should return true.")
	}
	card2.Value = 2
	if card1.IsEqual(card2) {
		t.Error("IsEqual should return false.")
	}
	card2.Color = "Heards"
	if card1.IsEqual(card2) {
		t.Error("IsEqual should return false.")
	}
	deck1 := NewDeck()
	deck2 := NewDeck()
	if !deck1.IsEqual(deck2) {
		t.Error("Is Equal should return true.")
	}
	deck1[1].Value = 10
	if deck1.IsEqual(deck2) {
		t.Error("Is Equal should return false.")
	}
}

func TestGetColor(t *testing.T) {
	for _, suit := range CardSuits {
		fmt.Printf("%s, %s", suit, getCardColor(suit))
		switch suit {
		case "Spades":
			if getCardColor(suit) != "Black" {
				t.Error("Spades should return Black.")
			}
		case "Hearts":
			if getCardColor(suit) != "Red" {
				t.Error("Hearts should return Red.")
			}
		case "Clubs":
			if getCardColor(suit) != "Black" {
				t.Error("Clubs should return Black.")
			}
		case "Diamonds":
			if getCardColor(suit) != "Red" {
				t.Error("Diamonds should return Red.")
			}
		}
	}
}

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
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

func TestRemoveCard(t *testing.T) {
	deck := NewDeck()
	card := deck.RemoveCard(0)
	if len(card) != 51 {
		t.Error("Only 1 card should be removed.")
	}
}

func TestRandomShuffle(t *testing.T) {
	deck := NewDeck()
	deck.RandomShuffle()
	if len(deck) != 52 {
		t.Error("Number of cards should stay the same.")
	}
	colorBlack := 0
	colorRed := 0
	spades := 0
	hearts := 0
	clubs := 0
	diamonds := 0
	value := 0
	for _, card := range deck {
		if card.Color == "Black" {
			colorBlack += 1
		} else if card.Color == "Red" {
			colorRed += 1
		}
		switch card.Suit {
		case "Spades":
			spades += 1
		case "Hearts":
			hearts += 1
		case "Clubs":
			clubs += 1
		case "Diamonds":
			diamonds += 1
		}
		value += card.Value
	}
	if colorBlack != 26 {
		t.Error("Incorrect number of black cards.")
	}
	if colorRed != 26 {
		t.Error("Incorrect number of red cards.")
	}
	if spades != 13 {
		t.Error("Incorrect number of spades.")
	}
	if hearts != 13 {
		t.Error("Incorrect number of hearts.")
	}
	if clubs != 13 {
		t.Error("Incorrect number of clubs.")
	}
	if diamonds != 13 {
		t.Error("Incorrect number of diamonds.")
	}
	if value != 364 {
		t.Error("Card values do not sum up to the correct value.")
	}
}

