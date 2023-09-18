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

func (g game) dealBoard() board {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.

	for i := 0; i < 7; i++ {
		cards := g.cards
		board := g.board
		// cards[:i+1].print()
		board[i][0] = cards[:i+1]
		board.printColumn(i)
		g.cards = cards[i+1:]
	}

	return board{}
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
	for _, column := range b {
		fmt.Print("[")
		for _, card := range column {
			card.print()
		}
		fmt.Println("]")
	}
}
