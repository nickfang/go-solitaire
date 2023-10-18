package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type game struct {
	cards            cards
	board            board
	stacks           stacks
	currentCardIndex int
	debug            bool
}

func newGame() game {
	return game{newDeck(), newBoard(), newStacks(), 0, false}
}

func (g *game) setDebug(onOff bool) {
	g.debug = onOff
	for i := range g.cards {
		g.cards[i].debug = onOff
	}
	for i, column := range g.board {
		for j := range column {
			g.board[i][j].debug = onOff
		}
	}
}

func (g *game) dealBoard() (cards, board, int) {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.
	cards := g.cards
	board := g.board
	currentCardIndex := 0

	for i := 0; i < NumColumns; i++ {
		for j := i; j < NumColumns; j++ {
			// fmt.Println(i, j, cards[j-i])
			board[j] = append(board[j], cards[j-i])
			// fmt.Println(board[j])
			if j == i {
				board[i][j].shown = true
			}
		}
		cards = cards[7-i:]
		// fmt.Println(cards)
	}
	cards[currentCardIndex].shown = true

	return cards, board, currentCardIndex
}

func (g *game) nextDeckCard() {
	g.cards[g.currentCardIndex].shown = false
	if g.currentCardIndex+3 > len(g.cards)-1 {
		g.currentCardIndex = 0
	} else {
		g.currentCardIndex += 3
	}
	g.cards[g.currentCardIndex].shown = true
}

// move the current card from the deck to a column
func (g *game) moveFromDeckToBoard(column int) {
	moves := g.getDeckMoves()
	if slices.Contains(moves, column) {
		g.cards[g.currentCardIndex].shown = true
		g.board[column] = append(g.board[column], g.cards[g.currentCardIndex])
		g.cards = g.cards.removeCard(g.currentCardIndex)
		if g.currentCardIndex > 0 {
			g.currentCardIndex = g.currentCardIndex - 1
		}
	} else {
		fmt.Println("Invalid move.")
	}
}

func (g *game) moveFromDeckToStacks() {
	moves := g.getStackMoves()
	currentCard := g.cards[g.currentCardIndex]
	var suitIndex int
	switch currentCard.suit {
	case CardSuits[0]:
		suitIndex = 0
	case CardSuits[1]:
		suitIndex = 1
	case CardSuits[2]:
		suitIndex = 2
	case CardSuits[3]:
		suitIndex = 3
	default:
	}
	if currentCard.value == moves[suitIndex] {
		g.stacks[suitIndex] = append(g.stacks[suitIndex], g.getCurrentCard())
		g.cards = g.cards.removeCard(g.currentCardIndex)
		if g.currentCardIndex > 0 {
			g.currentCardIndex = g.currentCardIndex - 1
		}
		return
	}
	fmt.Println("Invalid move.")
}

func (g *game) moveFromBoardToStacks(column int) {
	// move card from bottom of column to stacks
	moves := g.getStackMoves()
	lastIndex, lastCard := getLastCard(g.board[column])
	var suitIndex int
	switch lastCard.suit {
	case CardSuits[0]:
		suitIndex = 0
	case CardSuits[1]:
		suitIndex = 1
	case CardSuits[2]:
		suitIndex = 2
	case CardSuits[3]:
		suitIndex = 3
	default:
	}
	if lastCard.value == moves[suitIndex] {
		g.stacks[suitIndex] = append(g.stacks[suitIndex], lastCard)
		g.pruneColumn(column, lastIndex)
		columnLength := len(g.board[column])
		if columnLength > 0 && !g.board[column][columnLength-1].shown {
			g.board[column][len(g.board[column])-1].shown = true
		}
		return
	}
	fmt.Println("Invalid move.")
}

func (g *game) moveFromBoardToBoard(from int, to int) {
	// move cards from one column to another column
}

func (g *game) pruneColumn(column int, index int) []card {
	removed := g.board[column][index:]
	g.board[column] = g.board[column][:index]
	return removed
}

func (g game) getCurrentCard() card {
	return g.cards[g.currentCardIndex]
}

func checkMove(card card, toCard card) bool {
	if card.value == toCard.value-1 && card.color != toCard.color {
		return true
	}
	return false
}

func getLastCard(column []card) (int, card) {
	// turn an array into a slice so it's the right type.
	columnCopy := make([]card, len(column))
	copy(columnCopy, column[:])
	var lastIndex int
	var lastCard card
	for i, card := range columnCopy {
		lastIndex = i
		lastCard = card
		if card.value == 0 {
			if i == 0 {
				return i, card
			}
			return i - 1, columnCopy[i-1]
		}
	}
	return lastIndex, lastCard
}

// take the current deck card and return columns that are possible moves
// for the user the columns are 1 indexed instead of 0 indexed.
func (g game) getDeckMoves() []int {
	currentCard := g.getCurrentCard()
	moves := []int{}
	for index, column := range g.board {
		_, lastCard := getLastCard(column)
		if checkMove(currentCard, lastCard) {
			moves = append(moves, index)
		}
	}
	return moves
}

func (g game) getStackMoves() []int {
	moves := make([]int, 4)
	for i, stack := range g.stacks {
		moves[i] = len(stack) + 1
	}
	return moves
}

func (g game) displayCurrentCard() string {
	return g.getCurrentCard().displayMini
}

func (g game) displayCards() {
	fmt.Println(g.currentCardIndex, g.displayCurrentCard(), len(g.cards)-g.currentCardIndex)
}

func (g game) display() {
	g.stacks.display()
	g.board.display()
	g.displayCards()
	if g.debug {
		fmt.Println(g.getDeckMoves())
	}
}
