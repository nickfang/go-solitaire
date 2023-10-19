package board

import (
	"fmt"
	"solitaire/deck"
)

// using arrays since they are easier to initialize.
// TODO: maybe use slices and see how those work out.
type Board [][]deck.Card

const NumColumns = 7

func NewBoard() Board {
	board := Board{}
	for i := 0; i < NumColumns; i++ {
		board = append(board, []deck.Card{})
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

func (b Board) getColumnInfo(column int) {

}

func (b Board) Display() {
	displayBoard := [7][19]deck.Card{}
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
			displayBoard[x][y].Display()
		}
		fmt.Println()
	}
}
