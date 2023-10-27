package main

import "fmt"

type board [7][13]cards

type game struct {
	cards cards
	board board
}

func newGame() game {
	return game{newDeck(), newBoard()}
}

func (g game) dealBoard() (cards, board) {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.
	cards := g.cards
	board := g.board

	for i := 0; i < 7; i++ {
		board[i][0] = cards[:i+1]
		cards = cards[i+1:]
	}

	return cards, board
}

func newBoard() board {
	return board{}
}

func (b board) printColumn(column int) {
	for _, card := range b[column] {
		card.print()
	}
	fmt.Println()
}

func (b board) print() {
	for index, column := range b {
		fmt.Println("Column:", index+1)
		for _, card := range column {
			card.print()
		}
	}
}
