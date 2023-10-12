package main

import "fmt"

// using arrays since they are easier to initialize.
// TODO: maybe use slices and see how those work out.
type board [7][19]card

const NumColumns = 7

func newBoard() board {
	return board{}
}

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
