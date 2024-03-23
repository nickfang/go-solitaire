package board

import (
	"fmt"
	"solitaire/deck"
)

type Board [][]deck.Card

const NumColumns = 7

func (b1 Board) IsEqual(b2 Board) bool {
	for columnIndex, column := range b1 {
		if len(column) != len(b2[columnIndex]) {
			return false
		}
		for rowIndex, cell := range column {
			if !cell.IsEqual(b2[columnIndex][rowIndex]) {
				return false
			}
		}
	}
	return true
}

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

func (b Board) GetLastCard(column int) (int, deck.Card) {
	var lastIndex int
	var lastCard deck.Card
	for i, card := range b[column] {
		lastIndex = i
		lastCard = card
		if card.Value == 0 {
			if i == 0 {
				return i, card
			}
			return i - 1, b[column][i-1]
		}
	}
	return lastIndex, lastCard
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
