package main

import "fmt"

type board [19][7]card

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
		for j := 0; j < i+1; j++ {
			board[i][j] = cards[j]
			if j == i {
				board[i][j].shown = true
			}
		}
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

func (b board) display() {
	for y := 0; y < 19; y++ {
		for x := 0; x < 7; x++ {
			b[x][y].display()
		}
		fmt.Println()
	}
	// for i, row := range b {
	// 	for j, card := range row {
	// 		b[]
	// 		fmt.Print(i, j, ": ")
	// 		card.display()
	// 	}
	// 	// fmt.Println()
	// }
}
