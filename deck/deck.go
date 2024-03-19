package deck

import (
	"fmt"
	"math/rand"
	"time"
)

// var CardNumDisplay = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var CardNumDisplay = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var CardSuits = []string{"Spades", "Hearts", "Clubs", "Diamonds"}
var CardSuitsIcons = []string{"♠", "♥", "♣", "♦"}
var CardValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

type Card struct {
	Debug bool
	Shown bool
	Value int
	Suit  string
	Color string
	// display     string
	DisplayMini string
}

type Cards []Card

func getCardDisplay(value int, suit string) (string, error) {
	if value < 1 || value > 13 {
		return "", fmt.Errorf("invalid value: %d", value)
	}
	if value != 10 {
		return "  " + CardNumDisplay[value-1] + suit, nil
	}
	return " " + CardNumDisplay[value-1] + suit, nil
}

func getCardColor(suit string) string {
	if suit == "Spades" || suit == "Clubs" {
		return "Black"
	}
	return "Red"
}

func (c1 Card) IsEqual(c2 Card) bool {
	return c1.Suit == c2.Suit && c1.Value == c2.Value
}

func (d1 Cards) IsEqual(d2 Cards) bool {
	if len(d1) == 0 && len(d2) == 0 {
		return true
	}
	if len(d1) != len(d2) {
		return false // Decks have different number of cards
	}
	for i := range d1 {

		if !d1[i].IsEqual(d2[i]) {
			return false // Cards at the same position are different
		}
	}
	return true // All cards match in order
}

func NewCard(value int, suit string, shown bool) (Card, error) {
	suitIndex := -1
	for i, s := range CardSuits {
		if s == suit {
			suitIndex = i
		}
	}
	if suitIndex == -1 {
		return Card{}, fmt.Errorf("invalid suit: %s", suit)
	}
	displayMini, err := getCardDisplay(value, CardSuitsIcons[suitIndex])
	if err != nil {
		return Card{}, fmt.Errorf("invalid card display: %s", err)
	}
	return Card{false, shown, value, suit, getCardColor(suit), displayMini}, nil
}

func NewDeck() Cards {
	deck := Cards{}

	for _, suit := range CardSuits {
		for _, value := range CardValues {
			card, cardErr := NewCard(value, suit, false)
			if cardErr != nil {
				fmt.Printf("new deck not created: %s", cardErr.Error())
			}
			deck = append(
				deck,
				card,
			)
		}
	}

	return deck
}

func (d Cards) RemoveCard(cardIndex int) Cards {
	deck1 := d[:cardIndex]
	deck2 := d[cardIndex+1:]
	newDeck := append(deck1, deck2...)
	return newDeck
}

func (d Cards) RandomShuffle() {
	// todo: implement shuffle like hand shuffle
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d Cards) PerfectShuffle() {
	half1 := [26]Card(d[:len(d)/2])
	half2 := [26]Card(d[len(d)/2:])
	for i := 0; i < (len(d) / 2); i++ {
		d[i*2] = half1[i]
		d[(i*2)+1] = half2[i]
	}
}

func (d Cards) displayAll() {
	for _, card := range d {
		if card.DisplayMini == "null" {
			return
		} else if card.Shown || card.Debug {
			fmt.Print(card.DisplayMini)
		} else {
			fmt.Print("  * ")
		}
	}
}

func (c Card) Display() {
	if c.Value == 0 {
		fmt.Print("    ")
		return
	}
	if c.Shown || c.Debug {
		fmt.Print(c.DisplayMini)
	} else {
		fmt.Print("  * ")
	}
}

func (c Card) Print() {
	fmt.Printf("[V: %d, S: %s, C: %s]", c.Value, c.Suit, c.Color)
}

func (d Cards) Print() {
	for _, card := range d {
		card.Print()
	}
}
