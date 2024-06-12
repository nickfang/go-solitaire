package deck

import (
	"errors"
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

const DeckSize = 52

func getCardDisplay(value int, suit string) (string, error) {
	if value < 1 || value > 13 {
		return "", fmt.Errorf("invalid value: %d", value)
	}
	if suit != "♠" && suit != "♥" && suit != "♣" && suit != "♦" {
		return "", fmt.Errorf("invalid suit: %s", suit)
	}
	// add a space if the card value is 2 digits
	if value != 10 {
		return "  " + CardNumDisplay[value-1] + suit, nil
	}
	return " " + CardNumDisplay[value-1] + suit, nil
}

func getCardColor(suit string) string {
	if suit == "Spades" || suit == "Clubs" {
		return "Black"
	} else if suit == "Hearts" || suit == "Diamonds" {
		return "Red"
	}
	return ""
}

func (c1 Card) IsEqual(c2 Card) bool {
	return c1.Suit == c2.Suit && c1.Value == c2.Value
}

func (d1 Cards) IsEqual(d2 Cards) bool {
	if len(d1) == 0 && len(d2) == 0 {
		return true
	}
	if len(d1) != len(d2) {
		fmt.Println("lengths not equal", len(d1), len(d2))
		return false
	}
	for i := range d1 {
		if !d1[i].IsEqual(d2[i]) {
			fmt.Println("card not equal at index", i, d1[i], d2[i])
			return false
		}
	}
	return true
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
				return Cards{}
			}
			deck = append(
				deck,
				card,
			)
		}
	}

	return deck
}

func NewTestingDeck() (Cards, error) {
	deck := Cards{}
	for i := 13; i >= 7; i-- {
		for _, suit := range CardSuits {
			card, cardErr := NewCard(i, suit, true)
			if cardErr != nil {
				return nil, fmt.Errorf("new deck not created: %s", cardErr.Error())
			}
			deck = append(deck, card)
		}
	}
	for _, suit := range CardSuits {
		for _, value := range []int{3, 2, 1, 6, 5, 4} {
			card, cardErr := NewCard(value, suit, true)
			if cardErr != nil {
				return nil, fmt.Errorf("new deck not created: %s", cardErr.Error())
			}
			deck = append(deck, card)
		}
	}
	// for _, suit := range CardSuits {
	// 	for _, value := range []int{6, 5, 4} {
	// 		card, cardErr := NewCard(value, suit, true)
	// 		if cardErr != nil {
	// 			return nil, fmt.Errorf("new deck not created: %s", cardErr.Error())
	// 		}
	// 		deck = append(deck, card)
	// 	}
	// }
	return deck, nil
}

func (d *Cards) RemoveCard(cardIndex int) Card {
	length := len(*d)
	fmt.Println(length)
	card := (*d)[cardIndex]
	*d = append((*d)[:cardIndex], (*d)[cardIndex+1:length]...)
	return card
}

func (d *Cards) RandomShuffle() error {
	if d == nil {
		return errors.New("deck object required to shuffle cards")
	}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range *d {
		newPosition := r.Intn(len(*d) - 1)
		(*d)[i], (*d)[newPosition] = (*d)[newPosition], (*d)[i]
	}
	return nil
}

func (d *Cards) PerfectShuffle() {
	half1 := [26]Card((*d)[:len(*d)/2])
	half2 := [26]Card((*d)[len(*d)/2:])
	for i := 0; i < (len(*d) / 2); i++ {
		(*d)[i*2] = half1[i]
		(*d)[(i*2)+1] = half2[i]
	}
}

// CLI display
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

// Used for debugging
func (c Card) Print() {
	suit := " "
	if len(c.Suit) > 0 {
		if c.Color == "Red" {
			suit = "\033[31m" + c.Suit[:1] + "\033[0m"
		} else {
			suit = c.Suit[:1]
		}
	}
	fmt.Print("[")
	if c.Value >= 10 {
		fmt.Printf("%d", c.Value)
	} else if c.Value > 0 {
		fmt.Printf(" %d", c.Value)
	} else {
		fmt.Print("  ")
	}
	fmt.Printf("%s]", suit)
}

// Used for debugging
func (d Cards) Print() {
	for _, card := range d {
		card.Print()
	}
	fmt.Println()
}
