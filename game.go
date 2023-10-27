package main

import "fmt"

type board [7][19]card

type game struct {
	cards            cards
	board            board
	currentCardIndex int
}

func newGame() game {
	return game{newDeck(), newBoard(), 0}
}

func (g game) dealBoard() (cards, board) {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.
	cards := g.cards
	board := g.board

	for i := 0; i < 7; i++ {
		for j := i; j < 7; j++ {
			board[j][i] = cards[j-i]
			if j == i {
				board[i][j].shown = true
			}
		}
		// fmt.Println(board[i])
		cards = cards[7-i:]
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

func (b board) display() {
	for y := 0; y < 19; y++ {
		for x := 0; x < 7; x++ {
			b[x][y].display()
		}
		fmt.Println()
	}
}

func (g game) currentCard() card {
	return g.cards[g.currentCardIndex]
}

func (g game) getNextCard() int {
	if g.currentCardIndex+3 > len(g.cards)-1 {
		return 0
	}
	return g.currentCardIndex + 3
}
