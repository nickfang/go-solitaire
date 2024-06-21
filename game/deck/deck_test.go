package deck

import (
	"fmt"
	"testing"
)

func TestIsEqualCard(t *testing.T) {
	card1 := Card{
		false,
		false,
		1,
		"Spades",
		getCardColor("Spades"),
	}
	card2 := Card{
		false,
		false,
		1,
		"Spades",
		getCardColor("Spades"),
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
}
func TestIsEqualCards(t *testing.T) {
	emptyDeck1 := Cards{}
	emptyDeck2 := Cards{}
	if !emptyDeck1.IsEqual(emptyDeck2) {
		t.Error("Empty arrays should be equal.")
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

func TestNewCard(t *testing.T) {
	_, err := NewCard(1, "wrong", false)
	if err == nil {
		t.Error("invalid suit error show be returned.")
	}
	if err.Error() != "invalid suit: wrong" {
		t.Error("invalid suit error message is incorrect")
	}

	c1, err1 := NewCard(1, "Spades", false)
	if err1 != nil {
		t.Error("error while creating c1:", err1)
	}
	if c1.Suit != "Spades" || c1.Value != 1 || c1.Color != "Black" || c1.Shown != false {
		t.Error("c1 not created correctly", c1)
	}

	c2, err2 := NewCard(13, "Hearts", true)
	if err2 != nil {
		t.Error("error while creating c1:", err2)
	}
	if c2.Suit != "Hearts" || c2.Value != 13 || c2.Color != "Red" || c2.Shown != true {
		t.Error("c2 not created correctly", c2)
	}

	c3, err3 := NewCard(0, "Hearts", true)
	if !c3.IsEqual(Card{}) {
		t.Error("error while creating card should return default card.")
	}
	if err3.Error() != "invalid value: 0" {
		t.Error("Incorrect error message when value outside range.")
	}

	c4, err4 := NewCard(14, "Hearts", true)
	if !c4.IsEqual(Card{}) {
		t.Error("error while creating card should return default card.")
	}
	if err4.Error() != "invalid value: 14" {
		t.Error("Incorrect error message when value outside range.")
	}

	c5, err5 := NewCard(14, "Rocks", true)
	if !c5.IsEqual(Card{}) {
		t.Error("error while creating card should return default card.")
	}
	fmt.Print(err5.Error())
	if err5.Error() != "invalid suit: Rocks" {
		t.Error("Incorrect error message when value outside range.")
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

func TestNewTestingDeck(t *testing.T) {
	deck := NewDeck()
	err := deck.TestingShuffle()
	if err != nil {
		t.Error("Expected no error.")
	}
	// this order makes it easy to clear out the deck and board to test edge cases.
	deckValues := []int{13, 13, 13, 13, 12, 12, 12, 12, 11, 11, 11, 11, 10, 10, 10, 10, 9, 9, 9, 9, 8, 8, 8, 8, 7, 7, 7, 7, 3, 2, 1, 6, 5, 4, 3, 2, 1, 6, 5, 4, 3, 2, 1, 6, 5, 4, 3, 2, 1, 6, 5, 4}
	for i, card := range deck {
		if card.Value != deckValues[i] {
			t.Error("Card values are not correct.", card.Value, card)
		}
	}
}

func TestRemoveCard(t *testing.T) {
	deck := NewDeck()
	card, error := deck.RemoveCard(0)
	if error != nil {
		t.Error("Error should be nil.")
	}
	if card.Value != 1 {
		t.Error("Card value should be 1.")
	}
	if len(deck) != 51 {
		t.Error("Only 1 card should be removed.")
	}
	card, error = deck.RemoveCard(52)
	fmt.Println(error)
	if error.Error() != "error should be: invalid card index: 52" {
		t.Error()
	}
	if card.Value != 0 {
		t.Error("Card value should be 0.")
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
