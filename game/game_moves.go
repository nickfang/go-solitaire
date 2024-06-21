package game

import (
	"errors"
	"slices"
	// "solitaire/game/deck"
)

// move the current card from the deck to a column
func (g *Game) MoveFromDeckToBoard(column int) error {
	moves := g.GetDeckMoves()
	if !slices.Contains(moves, column) {
		return errors.New("invalid move")
	}
	g.Cards[g.CurrentCardIndex].Shown = true
	g.Board[column] = append(g.Board[column], g.Cards[g.CurrentCardIndex])
	g.Cards.RemoveCard(g.CurrentCardIndex)
	if g.CurrentCardIndex >= 0 {
		g.CurrentCardIndex = g.CurrentCardIndex - 1
	}
	return nil
}

func (g *Game) MoveFromDeckToStacks() error {
	if g.CurrentCardIndex < 0 {
		return errors.New("no cards in the deck")
	}
	suitIndex, validMove := g.GetStackMoves(g.Cards[g.CurrentCardIndex])
	if !validMove {
		return errors.New("invalid move")
	}
	movedCard, error := g.Cards.RemoveCard(g.CurrentCardIndex)
	if error != nil {
		return error
	}
	movedCard.Shown = true
	g.Stacks[suitIndex] = append(g.Stacks[suitIndex], movedCard)
	if g.CurrentCardIndex >= 0 {
		g.CurrentCardIndex = g.CurrentCardIndex - 1
	}
	return nil
}

func (g *Game) MoveFromBoardToStacks(column int) error {
	// move card from bottom of column to stacks
	lastIndex, lastCard := g.Board.GetLastCard(column)
	if lastIndex == -1 {
		return errors.New("nothing to move")
	}
	suitIndex, validMove := g.GetStackMoves(lastCard)
	if !validMove {
		return errors.New("invalid move")
	}
	g.Stacks[suitIndex] = append(g.Stacks[suitIndex], lastCard)
	g.pruneColumn(column, lastIndex)
	columnLength := len(g.Board[column])
	if columnLength > 0 && !g.Board[column][columnLength-1].Shown {
		g.Board[column][len(g.Board[column])-1].Shown = true
	}
	return nil
}

func (g *Game) MoveFromColumnToColumn(from int, to int) error {
	// move cards from one column to another column

	validIndex := -1
	moves := g.GetBoardMoves()

	for _, move := range moves {
		if move.FromColumn == from && move.ToColumn == to {
			validIndex = move.FromRow
			break
		}
	}

	if validIndex == -1 {
		return errors.New("invalid board move")
	}

	// remove the cards from the from column that were added to the to column
	removed := g.pruneColumn(from, validIndex)
	g.Board[to] = append(g.Board[to], removed...)
	if validIndex > 0 {
		g.Board[from][validIndex-1].Shown = true
	}

	// add all card from the king or the valid next card to the end of the from column to the to column
	return nil
}
