package game

import (
	"errors"
	"fmt"
	"strconv"

	"solitaire/game/board"
	"solitaire/game/deck"
	"solitaire/game/stacks"
)

type Game struct {
	Cards            deck.Cards
	Board            board.Board
	Stacks           stacks.Stacks
	CurrentCardIndex int
	Debug            bool
}

const NumCards = 52
const NumColumns = 7

func checkMove(card deck.Card, toCard deck.Card) bool {
	if card.Value == toCard.Value-1 && card.Color != toCard.Color {
		return true
	} else if card.Value == 13 && toCard.Value == 0 {
		return true
	}
	return false
}

func (g Game) getCurrentCard() (deck.Card, error) {
	if g.Cards == nil {
		return deck.Card{}, errors.New("deck not initialized")
	}
	if len(g.Cards) == 0 {
		return deck.Card{}, errors.New("no cards in the deck")
	}
	if g.CurrentCardIndex >= len(g.Cards) {
		return deck.Card{}, errors.New("current card index out of range")
	}

	return g.Cards[g.CurrentCardIndex], nil
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

func (g1 Game) IsEqual(g2 Game) bool {
	equalCards := g1.Cards.IsEqual(g2.Cards)
	equalStacks := g1.Stacks.IsEqual(g2.Stacks)
	equalBoard := g1.Board.IsEqual(g2.Board)
	equalCurrentCardIndex := g1.CurrentCardIndex == g2.CurrentCardIndex
	if g1.Debug {
		fmt.Printf("cards: %v\n", equalCards)
		fmt.Printf("stacks: %v\n", equalStacks)
		fmt.Printf("board: %v\n", equalBoard)
		fmt.Printf("index: %v\n", equalCurrentCardIndex)
	}
	if equalCards && equalStacks && equalBoard && equalCurrentCardIndex {
		return true
	}
	return false
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

func (g *Game) DealBoard() {
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
	g.Cards = cards
	g.Board = board
	g.CurrentCardIndex = currentCardIndex
}

func (g *Game) NextDeckCard() error {
	if len(g.Cards) == 0 {
		return errors.New("no more cards in the deck")
	}
	g.Cards[g.CurrentCardIndex].Shown = false
	if g.CurrentCardIndex+3 > len(g.Cards)-1 {
		g.CurrentCardIndex = 0
	} else {
		g.CurrentCardIndex += 3
	}
	g.Cards[g.CurrentCardIndex].Shown = true
	return nil
}

func (g Game) DeepCopy() Game {
	newState := Game{}

	// Deep Copy Board (assuming board.Board is [][]deck.Card)
	newState.Board = make(board.Board, len(g.Board))
	for i, row := range g.Board {
		newState.Board[i] = make([]deck.Card, len(row))
		copy(newState.Board[i], row)
	}

	// Deep Copy Cards
	newState.Cards = make(deck.Cards, len(g.Cards))
	copy(newState.Cards, g.Cards)

	// Deep Copy Stacks (assuming stacks.Stacks is [][]deck.Card)
	newState.Stacks = make(stacks.Stacks, len(g.Stacks))
	for i, suitStack := range g.Stacks {
		newState.Stacks[i] = make([]deck.Card, len(suitStack))
		copy(newState.Stacks[i], suitStack)
	}

	newState.CurrentCardIndex = g.CurrentCardIndex

	return newState
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
	currentCard, error := g.getCurrentCard()
	if error != nil {
		// no deck, no moves
		return []int{}
	}
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
	currentCard, error := g.getCurrentCard()
	if error != nil {
		// no current card, nothing to add to stackHints
	}
	_, validDeckMove := g.GetStackMoves(currentCard)
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
