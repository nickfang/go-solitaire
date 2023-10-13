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
}

func newGame() game {
	return game{newDeck(), newBoard(), newStacks(), 0}
}

func (g *game) setDebug(onOff bool) {
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
			board[j][i] = cards[j-i]
			if j == i {
				board[i][j].shown = true
			}
		}
		cards = cards[7-i:]
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
func (g *game) moveCurrentCard(column int) {
	moves := g.getDeckMoves()
	fmt.Println("moves", moves, column)
	if slices.Contains(moves, column) {
		lastIndex, _ := getLastCard(g.board[column])
		g.board[column][lastIndex+1] = g.cards[g.currentCardIndex]
		g.board[column][lastIndex+1].shown = true
		g.cards = g.cards.removeCard(g.currentCardIndex)
		if g.currentCardIndex > 0 {
			g.currentCardIndex = g.currentCardIndex - 1
		}
	} else {
		fmt.Println("Invalid move.")
	}
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

func getLastCard(column [19]card) (int, card) {
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
}
