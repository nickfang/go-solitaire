package main

import "fmt"

// using arrays since they are easier to initialize.
// TODO: maybe use slices and see how those work out.
type board [][]card

const NumColumns = 7

func newBoard() board {
	board := board{}
	for i := 0; i < NumColumns; i++ {
		board = append(board, []card{})
	}
	return board
}

// func (b board) pruneColumn(column int, index int) []card {
// 	removed := b[column][index:]
// 	columnCopy := make([]card, len(b[column]))
// 	copy(columnCopy, b[column][:])

// 	b[column] = columnCopy
// 	fmt.Println(removed)
// 	return removed
// }

func (b board) display() {
	displayBoard := [7][19]card{}
	maxLen := 0 // add a space so the board isn't cramped with the deck.
	for i, column := range b {
		if len(column) > maxLen {
			maxLen = len(column)
		}
		for j, card := range column {
			displayBoard[i][j] = card
		}
	}
	for y := 0; y < maxLen; y++ {
		for x := 0; x < 7; x++ {
			displayBoard[x][y].display()
		}
		fmt.Println()
	}
}
