package game

import (
	"fmt"

	"golang.org/x/exp/slices"
	"solitaire/board"
	"solitaire/deck"
	"solitaire/stacks"
)

type Game struct {
	Cards            deck.Cards
	Board            board.Board
	Stacks           stacks.Stacks
	CurrentCardIndex int
	Debug            bool
}

const NumColumns = 7

func NewGame() Game {
	return Game{deck.NewDeck(), board.NewBoard(), stacks.NewStacks(), 0, false}
}

func (g *Game) SetDebug(onOff bool) {
	g.Debug = onOff
	for i := range g.Cards {
		g.Cards[i].Debug = onOff
	}
	for i, column := range g.Board {
		for j := range column {
			g.Board[i][j].Debug = onOff
		}
	}
}

func (g *Game) DealBoard() (deck.Cards, board.Board, int) {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.
	cards := g.Cards
	board := g.Board
	currentCardIndex := 0

	for i := 0; i < NumColumns; i++ {
		for j := i; j < NumColumns; j++ {
			// fmt.Println(i, j, cards[j-i])
			board[j] = append(board[j], cards[j-i])
			// fmt.Println(board[j])
			if j == i {
				board[i][j].Shown = true
			}
		}
		cards = cards[7-i:]
		// fmt.Println(cards)
	}
	cards[currentCardIndex].Shown = true

	return cards, board, currentCardIndex
}

func (g *Game) NextDeckCard() {
	g.Cards[g.CurrentCardIndex].Shown = false
	if g.CurrentCardIndex+3 > len(g.Cards)-1 {
		g.CurrentCardIndex = 0
	} else {
		g.CurrentCardIndex += 3
	}
	g.Cards[g.CurrentCardIndex].Shown = true
}

// move the current card from the deck to a column
func (g *Game) MoveFromDeckToBoard(column int) {
	moves := g.GetDeckMoves()
	if slices.Contains(moves, column) {
		g.Cards[g.CurrentCardIndex].Shown = true
		g.Board[column] = append(g.Board[column], g.Cards[g.CurrentCardIndex])
		g.Cards = g.Cards.RemoveCard(g.CurrentCardIndex)
		if g.CurrentCardIndex > 0 {
			g.CurrentCardIndex = g.CurrentCardIndex - 1
		}
	} else {
		fmt.Println("Invalid move.")
	}
}

func (g *Game) MoveFromDeckToStacks() {
	moves := g.GetStackMoves()
	currentCard := g.Cards[g.CurrentCardIndex]
	var suitIndex int
	switch currentCard.Suit {
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
	if currentCard.Value == moves[suitIndex] {
		g.Stacks[suitIndex] = append(g.Stacks[suitIndex], g.getCurrentCard())
		g.Cards = g.Cards.RemoveCard(g.CurrentCardIndex)
		if g.CurrentCardIndex > 0 {
			g.CurrentCardIndex = g.CurrentCardIndex - 1
		}
		return
	}
	fmt.Println("Invalid move.")
}

func (g *Game) MoveFromBoardToStacks(column int) {
	// move card from bottom of column to stacks
	moves := g.GetStackMoves()
	lastIndex, lastCard := GetLastCard(g.Board[column])
	var suitIndex int
	switch lastCard.Suit {
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
	if lastCard.Value == moves[suitIndex] {
		g.Stacks[suitIndex] = append(g.Stacks[suitIndex], lastCard)
		g.pruneColumn(column, lastIndex)
		columnLength := len(g.Board[column])
		if columnLength > 0 && !g.Board[column][columnLength-1].Shown {
			g.Board[column][len(g.Board[column])-1].Shown = true
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
	if len(g.Board[to]) == 0 {
		validCard.Value = 13
	} else {
		lastCard := g.Board[to][len(g.Board[to])-1]
		fmt.Println(lastCard)
		if lastCard.Value != 1 {
			validCard.Value = lastCard.Value - 1
			if lastCard.Color == "Black" {
				validCard.Color = "Red"
			} else {
				validCard.Color = "Black"
			}
		}
	}
	visibleCards := []deck.Card{}
	for _, card := range g.Board[from] {
		if card.Shown {
			visibleCards = append(visibleCards, card)
		}
	}

	fmt.Println(visibleCards)
	if validCard.Value == 0 {
		fmt.Println("Invalid move from:", from, " to:", to)
		return
	}

	// add all card from the king or the valid next card to the end of the from column to the to column

	// remove the cards from the from column that were added to the to column

}

func (g *Game) pruneColumn(column int, index int) []deck.Card {
	removed := g.Board[column][index:]
	g.Board[column] = g.Board[column][:index]
	return removed
}

func (g Game) getCurrentCard() deck.Card {
	return g.Cards[g.CurrentCardIndex]
}

func CheckMove(card deck.Card, toCard deck.Card) bool {
	if card.Value == toCard.Value-1 && card.Color != toCard.Color {
		return true
	} else if card.Value == 13 && toCard.Value == 0 {
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
		if card.Value == 0 {
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
	for index, column := range g.Board {
		_, lastCard := GetLastCard(column)
		if CheckMove(currentCard, lastCard) {
			moves = append(moves, index)
		}
	}
	return moves
}

func (g Game) GetStackMoves() []int {
	moves := make([]int, 4)
	for i, stack := range g.Stacks {
		moves[i] = len(stack) + 1
	}
	return moves
}

func (g Game) DisplayCurrentCard() string {
	return g.getCurrentCard().DisplayMini
}

func (g Game) DisplayCards() {
	fmt.Println(g.CurrentCardIndex, g.DisplayCurrentCard(), len(g.Cards)-g.CurrentCardIndex)
}

func (g Game) Display() {
	g.Stacks.Display()
	g.Board.Display()
	g.DisplayCards()
	if g.Debug {
		fmt.Println(g.GetDeckMoves())
	}
}
