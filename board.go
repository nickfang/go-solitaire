package main

import "fmt"

type board [7][19]card

func (b board) display() {
	maxLen := 1 // add a space so the board isn't cramped with the deck.
	for _, column := range b {
		for index, card := range column {
			if card.value == 0 {
				maxLen = index
				break
			}
		}
	}

	for y := 0; y < maxLen; y++ {
		for x := 0; x < 7; x++ {
			b[x][y].display()
		}
		fmt.Println()
	}
}
