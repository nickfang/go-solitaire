package game

import (
	"fmt"
	"solitaire/game/board"
	"solitaire/game/deck"
	"solitaire/game/stacks"
)

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

	newState.FlipCount = g.FlipCount

	return newState
}

func (g1 Game) IsEqual(g2 Game) bool {
	equalCards := g1.Cards.IsEqual(g2.Cards)
	equalStacks := g1.Stacks.IsEqual(g2.Stacks)
	equalBoard := g1.Board.IsEqual(g2.Board)
	equalCurrentCardIndex := g1.CurrentCardIndex == g2.CurrentCardIndex
	equalFlipCount := g1.FlipCount == g2.FlipCount
	if g1.Debug {
		fmt.Printf("cards: %v\n", equalCards)
		fmt.Printf("stacks: %v\n", equalStacks)
		fmt.Printf("board: %v\n", equalBoard)
		fmt.Printf("index: %v\n", equalCurrentCardIndex)
		fmt.Printf("flip count: %v\n", equalFlipCount)
	}
	if equalCards && equalStacks && equalBoard && equalCurrentCardIndex && equalFlipCount {
		return true
	}
	return false
}

func (g *Game) pruneColumn(column int, index int) []deck.Card {
	removed := g.Board[column][index:]
	g.Board[column] = g.Board[column][:index]
	return removed
}

func (g Game) CheckGame() error {
	numCards := 0
	for _, stack := range g.Stacks {
		numCards += len(stack)
	}
	for _, column := range g.Board {
		numCards += len(column)
	}
	numCards += len(g.Cards)
	if numCards != 52 {
		return fmt.Errorf("invalid game state: %d cards", numCards)
	}
	return nil
}
