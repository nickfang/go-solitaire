package main

import "fmt"

type game struct {
	cards            cards
	board            board
	currentCardIndex int
}

func newGame() game {
	return game{newDeck(), newBoard(), 0}
}

func (g game) setDebug() {
	for _, card := range g.cards {
		card.debug = true
	}
}

func (g game) dealBoard() (cards, board, int) {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.
	cards := g.cards
	board := g.board
	currentCardIndex := 0

	for i := 0; i < 7; i++ {
		for j := i; j < 7; j++ {
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

func newBoard() board {
	return board{}
}

func (g game) currentCard() card {
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

func (g game) displayCards() {
	fmt.Println(g.currentCardIndex, g.cards[g.currentCardIndex].displayMini, len(g.cards)-g.currentCardIndex)
}

func (g game) displayBoard() {
	maxLen := 1 // add a space so the board isn't cramped with the deck.
	for _, column := range g.board {
		for index, card := range column {
			if card.value == 0 {
				maxLen = index
				break
			}
		}
	}

	for y := 0; y < maxLen; y++ {
		for x := 0; x < 7; x++ {
			g.board[x][y].display()
		}
		fmt.Println()
	}
}
