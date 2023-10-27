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

func (g game) getCurrentCard() card {
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

func checkMove(card card, toCard card) bool {
	if card.value == toCard.value-1 && card.color != toCard.color {
		return true
	}
	return false
}

func getLastCard(cards []card) (int, card) {
	var lastIndex int
	var lastCard card
	for i, card := range cards {
		lastIndex = i
		lastCard = card
		if card.value == 0 {
			if i == 0 {
				return i, card
			}
			return i - 1, cards[i-1]
		}
	}
	return lastIndex, lastCard
}

// take the current deck card and return columns that are possible moves
func (g game) getDeckMoves() []int {
	currentCard := g.getCurrentCard()
	moves := []int{}
	for index, column := range g.board {
		columnCopy := make([]card, len(column))
		copy(columnCopy, column[:])
		_, lastCard := getLastCard(columnCopy)
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

func (g game) displayBoard() {
	maxLen := 0
	for _, column := range g.board {
		// turn an array into a slice so it's the right type.
		columnCopy := make([]card, len(column))
		copy(columnCopy, column[:])
		index, _ := getLastCard(columnCopy)
		if maxLen < index+1 {
			maxLen = index + 1
		}
	}

	for y := 0; y < maxLen; y++ {
		for x := 0; x < 7; x++ {
			g.board[x][y].display()
		}
		fmt.Println()
	}
}
