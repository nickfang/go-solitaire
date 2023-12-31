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

//	func getCardDisplay(value int, suit string) string {
//		return CardNumDisplay[value-1] + " of " + suit
//	}
func getCardDisplay(value int, suit string) string {
	if value != 10 {
		return "  " + CardNumDisplay[value-1] + suit
	}
	return " " + CardNumDisplay[value-1] + suit
}
func getCardColor(suit string) string {
	if suit == "Spades" || suit == "Clubs" {
		return "Black"
	}
	return "Red"
}

func NewDeck() Cards {
	deck := Cards{}

	for index, suit := range CardSuits {
		for _, value := range CardValues {
			deck = append(
				deck,
				Card{
					false,
					false,
					value,
					suit,
					getCardColor(suit),
					getCardDisplay(value, CardSuitsIcons[index]),
				},
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

func (d Cards) perfectShuffle() {
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
