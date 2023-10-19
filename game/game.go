package game

import (
	"fmt"

	"golang.org/x/exp/slices"
	"solitaire/board"
	"solitaire/deck"
	"solitaire/stacks"
)

type Game struct {
	cards            deck.Cards
	board            board.Board
	stacks           stacks.Stacks
	currentCardIndex int
	debug            bool
}

func NewGame() Game {
	return Game{deck.NewDeck(), board.NewBoard(), stacks.NewStacks(), 0, false}
}

func (g *Game) setDebug(onOff bool) {
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

func (g *Game) DealBoard() (deck.Cards, board.Board, int) {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.
	cards := g.cards
	board := g.board
	currentCardIndex := 0

	for i := 0; i < board.NumColumns; i++ {
		for j := i; j < board.NumColumns; j++ {
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

func (g *Game) NextDeckCard() {
	g.cards[g.currentCardIndex].shown = false
	if g.currentCardIndex+3 > len(g.cards)-1 {
		g.currentCardIndex = 0
	} else {
		g.currentCardIndex += 3
	}
	g.cards[g.currentCardIndex].shown = true
}

// move the current card from the deck to a column
func (g *Game) MoveFromDeckToBoard(column int) {
	moves := g.GetDeckMoves()
	if slices.Contains(moves, column) {
		g.cards[g.currentCardIndex].shown = true
		g.board[column] = append(g.board[column], g.cards[g.currentCardIndex])
		g.cards = g.cards.RemoveCard(g.currentCardIndex)
		if g.currentCardIndex > 0 {
			g.currentCardIndex = g.currentCardIndex - 1
		}
	} else {
		fmt.Println("Invalid move.")
	}
}

func (g *Game) MoveFromDeckToStacks() {
	moves := g.GetStackMoves()
	currentCard := g.cards[g.currentCardIndex]
	var suitIndex int
	switch currentCard.suit {
	case deck.CardSuits[0]:
		suitIndex = 0
	case deck.CardSuits[1]:
		suitIndex = 1
	case deck.CardSuits[2]:
		suitIndex = 2
	case deck.CardSuits[3]:
		suitIndex = 3
	default:
	}
	if currentCard.value == moves[suitIndex] {
		g.stacks[suitIndex] = append(g.stacks[suitIndex], g.getCurrentCard())
		g.cards = g.cards.RemoveCard(g.currentCardIndex)
		if g.currentCardIndex > 0 {
			g.currentCardIndex = g.currentCardIndex - 1
		}
		return
	}
	fmt.Println("Invalid move.")
}

func (g *Game) MoveFromBoardToStacks(column int) {
	// move card from bottom of column to stacks
	moves := g.GetStackMoves()
	lastIndex, lastCard := GetLastCard(g.board[column])
	var suitIndex int
	switch lastCard.suit {
	case deck.CardSuits[0]:
		suitIndex = 0
	case deck.CardSuits[1]:
		suitIndex = 1
	case deck.CardSuits[2]:
		suitIndex = 2
	case deck.CardSuits[3]:
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

func (g *Game) MoveFromColumnToColumn(from int, to int) {
	// move cards from one column to another column
	// check if this is a valid move
	// one of the cards showing in the from column can be put on the last card of the to column
	// If the to column is empty, make sure a king is the first card of the from column.
	var validCard deck.Card
	if len(g.board[to]) == 0 {
		validCard.value = 13
	} else {
		lastCard := g.board[to][len(g.board[to])-1]
		fmt.Println(lastCard)
		if lastCard.value != 1 {
			validCard.value = lastCard.value - 1
			if lastCard.color == "Black" {
				validCard.color = "Red"
			} else {
				validCard.color = "Black"
			}
		}
	}
	visibleCards := []deck.Card{}
	for _, card := range g.board[from] {
		if card.shown {
			visibleCards = append(visibleCards, card)
		}
	}

	fmt.Println(visibleCards)
	if validCard.value == 0 {
		fmt.Println("Invalid move from:", from, " to:", to)
		return
	}

	// add all card from the king or the valid next card to the end of the from column to the to column

	// remove the cards from the from column that were added to the to column

}

func (g *Game) pruneColumn(column int, index int) []deck.Card {
	removed := g.board[column][index:]
	g.board[column] = g.board[column][:index]
	return removed
}

func (g Game) getCurrentCard() deck.Card {
	return g.cards[g.currentCardIndex]
}

func CheckMove(card deck.Card, toCard deck.Card) bool {
	if card.value == toCard.value-1 && card.color != toCard.color {
		return true
	} else if card.value == 13 && toCard.value == 0 {
	}
	return false
}

func GetLastCard(column []deck.Card) (int, deck.Card) {
	// turn an array into a slice so it's the right type.
	columnCopy := make([]deck.Card, len(column))
	copy(columnCopy, column[:])
	var lastIndex int
	var lastCard deck.Card
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
func (g Game) GetDeckMoves() []int {
	currentCard := g.getCurrentCard()
	moves := []int{}
	for index, column := range g.board {
		_, lastCard := GetLastCard(column)
		if CheckMove(currentCard, lastCard) {
			moves = append(moves, index)
		}
	}
	return moves
}

func (g Game) GetStackMoves() []int {
	moves := make([]int, 4)
	for i, stack := range g.stacks {
		moves[i] = len(stack) + 1
	}
	return moves
}

func (g Game) DisplayCurrentCard() string {
	return g.getCurrentCard().displayMini
}

func (g Game) DisplayCards() {
	fmt.Println(g.currentCardIndex, g.DisplayCurrentCard(), len(g.cards)-g.currentCardIndex)
}

func (g Game) Display() {
	g.stacks.Display()
	g.board.Display()
	g.DisplayCards()
	if g.debug {
		fmt.Println(g.GetDeckMoves())
	}
}
