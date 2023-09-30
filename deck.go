package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var cardNumDisplay = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var cardNumDisplay = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

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
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardSuitsIcons := []string{"♠", "♥", "♣", "♦"}
	cardValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	for index, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(
				deck,
				card{
					false,
					false,
					value,
					suit,
					getCardColor(suit),
					getCardDisplay(value, cardSuitsIcons[index]),
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

func (c cards) randomShuffle() {
	// todo: implement shuffle like hand shuffle
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range c {
		newPosition := r.Intn(len(c) - 1)
		c[i], c[newPosition] = c[newPosition], c[i]
	}
}

func (c cards) perfectShuffle() {
	half1 := [26]card(c[:len(c)/2])
	half2 := [26]card(c[len(c)/2:])
	for i := 0; i < (len(c) / 2); i++ {
		c[i*2] = half1[i]
		c[(i*2)+1] = half2[i]
	}
}

func (c cards) displayAll() {
	for _, card := range c {
		if card.displayMini == "null" {
			return
		} else if card.shown {
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
	if c.debug {
		fmt.Print(c.displayMini)
	} else if c.shown {
		fmt.Print(c.displayMini)
	} else {
		fmt.Print("  * ")
	}
}
