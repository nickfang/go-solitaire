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

func setDebug(g game) game {
	cards := g.cards
	for _, card := range cards {
		card.debug = true
	}
	return g
}

func (g game) dealBoard() (cards, board) {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.
	cards := g.cards
	board := g.board
	// currentCardIndex := 0

	for i := 0; i < 7; i++ {
		for j := i; j < 7; j++ {
			board[j][i] = cards[j-i]
			if j == i {
				board[i][j].shown = true
			}
		}
		cards = cards[7-i:]
	}
	// cards[currentCardIndex].shown = true

	return cards, board
}

func newBoard() board {
	return board{}
}

// func (b board) printColumn(column int) {
// 	for _, card := range b[column] {
// 		card.print()
// 	}
// 	fmt.Println()
// }

// func (b board) print() {
// 	for index, column := range b {
// 		fmt.Println("Column:", index+1)
// 		for _, card := range column {
// 			card.print()
// 		}
// 	}
// }

func (b board) display() {
	for y := 0; y < 19; y++ {
		for x := 0; x < 7; x++ {
			b[x][y].display()
		}
		fmt.Println()
	}
}

func (g game) currentCard() card {
	// g.cards[g.currentCardIndex].shown = true
	return g.cards[g.currentCardIndex]
}

func (g game) getNextCard() game {
	g.cards[g.currentCardIndex].shown = false
	if g.currentCardIndex+3 > len(g.cards)-1 {
		g.currentCardIndex = 0
	} else {
		g.currentCardIndex += 3
	}
	g.cards[g.currentCardIndex].shown = true
	return g
}
