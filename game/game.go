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
	Id               string
	Cards            deck.Cards
	Board            board.Board
	Stacks           stacks.Stacks
	CurrentCardIndex int
	FlipCount        int
	Debug            bool
}

type BoardMove struct {
	FromColumn int
	FromRow    int
	ToColumn   int
}

const DefaultFlipCount = 3

func checkMove(card deck.Card, toCard deck.Card) bool {
	if card.Value == toCard.Value-1 && card.Color != toCard.Color {
		return true
	} else if card.Value == 13 && toCard.Value == 0 {
		// special case, if a column is empty, we return a default card which has a value of 0
		return true
	}
	return false
}

func (g Game) GetCurrentCard() (deck.Card, error) {
	if g.Cards == nil {
		return deck.Card{}, errors.New("deck not initialized")
	}
	if len(g.Cards) == 0 {
		return deck.Card{}, errors.New("no cards in the deck")
	}
	if g.CurrentCardIndex >= len(g.Cards) {
		return deck.Card{}, errors.New("current card index out of range")
	}
	if g.CurrentCardIndex == -1 {
		return deck.Card{}, nil
	}
	return g.Cards[g.CurrentCardIndex], nil
}

/* Exported Functions */

func NewGame(id string) Game {
	newId := id
	if id == "" {
		newId = "main"
	}
	return Game{newId, deck.NewDeck(), board.NewBoard(), stacks.NewStacks(), DefaultFlipCount - 1, DefaultFlipCount, false}
}

func (g *Game) Reset() {
	g.Cards = deck.NewDeck()
	g.Board = board.NewBoard()
	g.Stacks = stacks.NewStacks()
	g.FlipCount = DefaultFlipCount
	g.CurrentCardIndex = DefaultFlipCount - 1
}

func (g *Game) DealBoard() {
	// a board of solitare is 7 columns of cards
	// the first column has 1 card, the second has 2, etc.
	gameCards := g.Cards
	gameBoard := g.Board
	currentCardIndex := DefaultFlipCount - 1

	for i := 0; i < board.NumColumns; i++ {
		for j := i; j < board.NumColumns; j++ {
			gameBoard[j] = append(gameBoard[j], gameCards[j-i])
			if j == i {
				gameBoard[i][j].Shown = true
			}
		}
		gameCards = gameCards[7-i:]
	}
	gameCards[currentCardIndex].Shown = true
	g.Cards = gameCards
	g.Board = gameBoard
	g.CurrentCardIndex = currentCardIndex
}

func (g *Game) NextDeckCard() error {
	deckLength := len(g.Cards)
	if deckLength == 0 {
		return errors.New("no more cards in the deck")
	}
	// if there is a current card, hide it
	if g.CurrentCardIndex >= 0 {
		g.Cards[g.CurrentCardIndex].Shown = false
	}
	// if the next card is out of bounds, set the current card back to the beginning
	if g.CurrentCardIndex == -1 || g.CurrentCardIndex+g.FlipCount > deckLength-1 {
		if 2 < deckLength-1 {
			g.CurrentCardIndex = g.FlipCount - 1
		} else {
			g.CurrentCardIndex = deckLength - 1
		}
	} else {
		g.CurrentCardIndex += g.FlipCount
	}
	g.Cards[g.CurrentCardIndex].Shown = true
	return nil
}

func (g *Game) UpdateState(gameState Game) error {
	error := gameState.CheckGame()
	if error != nil {
		return errors.New("game state is nil")
	}
	g.Cards = gameState.Cards
	g.Board = gameState.Board
	g.Stacks = gameState.Stacks
	g.CurrentCardIndex = gameState.CurrentCardIndex
	g.FlipCount = gameState.FlipCount
	return nil
}

func (g *Game) SetFlipCount(flipCount int) error {
	if flipCount != 1 && flipCount != 3 {
		return errors.New("flip count must be 1 or 3")
	}
	g.FlipCount = flipCount
	g.CurrentCardIndex = flipCount - 1
	return nil
}

func (g Game) IsFinished() bool {
	for _, stack := range g.Stacks {
		if len(stack) != 13 {
			return false
		}
	}
	return true
}

// take the current deck card and return columns that are possible moves
// for the user the columns are 1 indexed instead of 0 indexed.
func (g Game) GetDeckMoves() []int {
	moves := []int{}
	currentCard, error := g.GetCurrentCard()
	if error != nil {
		// no deck, no moves
		return moves
	}
	for index := range g.Board {
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

func (g Game) GetBoardMoves() []BoardMove {
	moves := []BoardMove{}
	lastCards := deck.Cards{}
	for i := range g.Board {
		_, lastCard := g.Board.GetLastCard(i)
		lastCards = append(lastCards, lastCard)
	}
	for i, column := range g.Board {
		for j, card := range column {
			if !card.Shown {
				continue
			}
			// see if current shown card can be moved to any of the last cards
			for k, lastCard := range lastCards {
				if checkMove(card, lastCard) {
					boardMove := BoardMove{i, j, k}
					moves = append(moves, boardMove)
				}
			}
		}
	}
	return moves
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
	currentCard, error := g.GetCurrentCard()
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

func (g Game) GetBoardHints() []string {
	hints := []string{}
	for _, move := range g.GetBoardMoves() {
		moveStr := strconv.FormatInt(int64(move.FromColumn), 10) + strconv.FormatInt(int64(move.ToColumn), 10)
		hints = append(hints, moveStr)
	}
	return hints
}
