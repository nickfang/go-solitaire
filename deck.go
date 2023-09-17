package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cardNumDisplay = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

type card struct {
	value   int
	suit    string
	color   string
	display string
}

type cards []card

type game struct {
	cards cards
}

func getCardDisplay(value int, suit string) string {
	return cardNumDisplay[value-1] + " of " + suit
}
func getCardColor(suit string) string {
	if suit == "Spades" || suit == "Clubs" {
		return "Black"
	}
	return "Red"
}

func newGame() game {
	return game{newDeck()}
}

func newDeck() cards {
	deck := cards{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, card{value, suit, getCardColor(suit), getCardDisplay(value, suit)})
		}
	}

	return deck
}

func (g game) randomShuffle() {
	// todo: implement shuffle like hand shuffle
	c := g.cards
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range c {
		newPosition := r.Intn(len(c) - 1)
		c[i], c[newPosition] = c[newPosition], c[i]
	}
}

func (g game) perfectShuffle() {
	c := g.cards
	half1 := [26]card(c[:len(c)/2])
	half2 := [26]card(c[len(c)/2:])
	for i := 0; i < (len(c) / 2); i++ {
		c[i*2] = half1[i]
		c[(i*2)+1] = half2[i]
	}
}

func (g game) printCards() {
	for _, card := range g.cards {
		fmt.Println(card.display)
	}
}
