package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var cardNumDisplay = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var cardNumDisplay = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var CardSuits = []string{"Spades", "Hearts", "Clubs", "Diamonds"}
var CardSuitsIcons = []string{"♠", "♥", "♣", "♦"}
var CardValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

type card struct {
	debug bool
	shown bool
	value int
	suit  string
	color string
	// display     string
	displayMini string
}

type cards []card

//	func getCardDisplay(value int, suit string) string {
//		return cardNumDisplay[value-1] + " of " + suit
//	}
func getCardDisplay(value int, suit string) string {
	if value != 10 {
		return "  " + cardNumDisplay[value-1] + suit
	}
	return " " + cardNumDisplay[value-1] + suit
}
func getCardColor(suit string) string {
	if suit == "Spades" || suit == "Clubs" {
		return "Black"
	}
	return "Red"
}

func newDeck() cards {
	deck := cards{}

	for index, suit := range CardSuits {
		for _, value := range CardValues {
			deck = append(
				deck,
				card{
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

func (d cards) removeCard(cardIndex int) cards {
	deck1 := d[:cardIndex]
	deck2 := d[cardIndex+1:]
	newDeck := append(deck1, deck2...)
	return newDeck
}

func (d cards) randomShuffle() {
	// todo: implement shuffle like hand shuffle
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d cards) perfectShuffle() {
	half1 := [26]card(d[:len(d)/2])
	half2 := [26]card(d[len(d)/2:])
	for i := 0; i < (len(d) / 2); i++ {
		d[i*2] = half1[i]
		d[(i*2)+1] = half2[i]
	}
}

func (d cards) displayAll() {
	for _, card := range d {
		if card.displayMini == "null" {
			return
		} else if card.shown || card.debug {
			fmt.Print(card.displayMini)
		} else {
			fmt.Print("  * ")
		}
	}
}

func (c card) display() {
	if c.value == 0 {
		fmt.Print("    ")
		return
	}
	if c.debug || c.shown {
		fmt.Print(c.displayMini)
	} else {
		fmt.Print("  * ")
	}
}
