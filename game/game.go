package game

import (
	"fmt"
	"strconv"

	"solitaire/board"
	"solitaire/deck"
	"solitaire/stacks"

	"golang.org/x/exp/slices"
)

type Game struct {
	Cards            deck.Cards
	Board            board.Board
	Stacks           stacks.Stacks
	CurrentCardIndex int
	Debug            bool
}

const NumColumns = 7

func checkMove(card deck.Card, toCard deck.Card) bool {
	if card.Value == toCard.Value-1 && card.Color != toCard.Color {
		return true
	} else if card.Value == 13 && toCard.Value == 0 {
		return true
	}
	return false
}

func (g Game) getCurrentCard() deck.Card {
	return g.Cards[g.CurrentCardIndex]
}

func (g *Game) pruneColumn(column int, index int) []deck.Card {
	removed := g.Board[column][index:]
	g.Board[column] = g.Board[column][:index]
	return removed
}


/* Exported Functions */

func NewGame() Game {
	return Game{deck.NewDeck(), board.NewBoard(), stacks.NewStacks(), 0, false}
}

func (g *Game) Reset() {
	g.Cards = deck.NewDeck()
	g.Board = board.NewBoard()
	g.Stacks = stacks.NewStacks()
	g.CurrentCardIndex = 0
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
	currentCard := g.Cards[g.CurrentCardIndex]
	suitIndex, validMove := g.GetStackMoves(currentCard)
	if validMove {
		currentCard.Shown = true
		g.Stacks[suitIndex] = append(g.Stacks[suitIndex], currentCard)
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
	lastIndex, lastCard := g.Board.GetLastCard(column)
	suitIndex, validMove := g.GetStackMoves(lastCard)
	if validMove {
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

	// get the last card of the to column and figure out what value and color can be placed on top of it.
	// if it is an empty column, the card must be a king of any color.
	var validCard deck.Card
	if len(g.Board[to]) == 0 {
		validCard.Value = 13
	} else {
		lastCard := g.Board[to][len(g.Board[to])-1]
		if lastCard.Value != 1 {
			validCard.Value = lastCard.Value - 1
			if lastCard.Color == "Black" {
				validCard.Color = "Red"
			} else {
				validCard.Color = "Black"
			}
		}
	}

	if validCard.Value == 0 {
		fmt.Println("Invalid Move.  Cannot place a card on an ace.")
		return
	}

	validIndex := -1
	for index, card := range g.Board[from] {
		if card.Shown {
			if validCard.Value == 13 && card.Value == 13 {
				validIndex = index
			} else if card.Value == validCard.Value && card.Color == validCard.Color {
				validIndex = index
			}
		}
	}

	if validIndex == -1 {
		fmt.Println("Invalid Move. No valid cards to move.")
		return
	}

	// remove the cards from the from column that were added to the to column
	removed := g.pruneColumn(from, validIndex)
	g.Board[to] = append(g.Board[to], removed...)
	if validIndex > 0 {
		g.Board[from][validIndex-1].Shown = true
	}

	// add all card from the king or the valid next card to the end of the from column to the to column

}

func (g *Game) SetState(gameState Game) {
	g.Cards = gameState.Cards
	g.Board = gameState.Board
	g.Stacks = gameState.Stacks
	g.CurrentCardIndex = gameState.CurrentCardIndex
}


// func GetLastCard(column []deck.Card) (int, deck.Card) {
// 	// turn an array into a slice so it's the right type.
// 	columnCopy := make([]deck.Card, len(column))
// 	copy(columnCopy, column[:])
// 	var lastIndex int
// 	var lastCard deck.Card
// 	for i, card := range columnCopy {
// 		lastIndex = i
// 		lastCard = card
// 		if card.Value == 0 {
// 			if i == 0 {
// 				return i, card
// 			}
// 			return i - 1, columnCopy[i-1]
// 		}
// 	}
// 	return lastIndex, lastCard
// }

// take the current deck card and return columns that are possible moves
// for the user the columns are 1 indexed instead of 0 indexed.


func (g Game) GetDeckMoves() []int {
	currentCard := g.getCurrentCard()
	moves := []int{}
	for index, _ := range g.Board {
		_, lastCard := g.Board.GetLastCard(index)
		if checkMove(currentCard, lastCard) {
			moves = append(moves, index)
		}
	}
	return moves
}

func (g Game) GetStackMoves(card deck.Card) (int, bool) {
	moves := make([]int, 4)
	for i, stack := range g.Stacks {
		moves[i] = len(stack) + 1
	}

	var suitIndex int
	switch card.Suit {
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
	if card.Value == moves[suitIndex] {
		return suitIndex, true
	}
	return suitIndex, false
}

func (g Game) GetDeckHints() []string {
	deckHints := []string{}
	for _, move := range g.GetDeckMoves() {
		fmt.Println(move)
		moveStr := "d" + strconv.FormatInt(int64(move), 10)
		deckHints = append(deckHints, moveStr)
	}
	return deckHints
}

func (g Game) GetStackHints() []string {
	stackHints := []string{}
	// check deck first
	_, validDeckMove := g.GetStackMoves(g.getCurrentCard())
	if validDeckMove {
		stackHints = append(stackHints, "ds")
	}
	for i := 0; i < len(g.Board); i++ {
		_, lastCard := g.Board.GetLastCard(i)
		_, validBoardMove := g.GetStackMoves(lastCard)
		if validBoardMove {
			moveStr := strconv.FormatInt(int64(i), 10) + "s"
			stackHints = append(stackHints, moveStr)
		}
	}
	return stackHints
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

func (g Game) DisplayHints() {
	deckMoves := g.GetDeckHints()
	stackMoves := g.GetStackHints()
	fmt.Println("Moves:", append(deckMoves, stackMoves...))
}
