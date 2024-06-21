package main

import (
	"errors"
	"fmt"
	"strconv"

	"solitaire/game"
	"solitaire/game/gamestates"

	"golang.org/x/exp/slices"
)

func NextCard(g *game.Game, gs *gamestates.GameStates) error {
	nextErr := g.NextDeckCard()
	if nextErr != nil {
		return nextErr
	}
	gs.SaveState(*g)
	return nil
}

func ResetGame(g *game.Game, gs *gamestates.GameStates) error {
	g.Reset()
	shuffleError := g.Cards.RandomShuffle()
	if shuffleError != nil {
		return shuffleError
	}
	g.DealBoard()
	gs.Reset()
	gs.SaveState(*g)
	return nil
}

func ShowHints(g *game.Game) {
	moves := g.GetDeckHints()
	moves = append(moves, g.GetStackHints()...)
	moves = append(moves, g.GetBoardHints()...)
	fmt.Println("Moves:", moves)
}

func Undo(g *game.Game, gs *gamestates.GameStates) error {
	if len(gs.States) <= 1 {
		return errors.New("no moves to undo")
	}
	lastGameState := gs.Undo()
	error := g.UpdateState(lastGameState)
	if error != nil {
		return error
	}
	return nil
}

func MoveDeckToBoard(input1 string, g *game.Game, gs *gamestates.GameStates) error {
	if !slices.Contains(ValidColumns, input1) {
		return errors.New("invalid column: " + input1)
	}
	columnIndex, _ := strconv.ParseInt(input1, 10, 32)
	error := g.MoveFromDeckToBoard(int(columnIndex - 1))
	if error != nil {
		return error
	}
	gs.SaveState(*g)
	return nil
}

func MoveDeckToStacks(g *game.Game, gs *gamestates.GameStates) error {
	error := g.MoveFromDeckToStacks()
	if error != nil {
		return error
	}
	gs.SaveState(*g)
	return nil
}

func MoveBoardToStacks(input0 string, g *game.Game, gs *gamestates.GameStates) error {
	columnIndex, _ := strconv.ParseInt(input0, 10, 32)
	error := g.MoveFromBoardToStacks(int(columnIndex - 1))
	if error != nil {
		return error
	}
	gs.SaveState(*g)
	return nil
}

func MoveColumnToColumn(input0, input1 string, g *game.Game, gs *gamestates.GameStates) error {
	if (!slices.Contains(ValidColumns, input0) || !slices.Contains(ValidColumns, input1)) || input0 == input1 {
		return errors.New("invalid move input")
	}

	fromColumn, _ := strconv.ParseInt(input0, 10, 32)
	toColumn, _ := strconv.ParseInt(input1, 10, 32)
	error := g.MoveFromColumnToColumn(int(fromColumn-1), int(toColumn-1))
	if error != nil {
		return error
	}

	gs.SaveState(*g)
	return nil
}

func DealTest(g *game.Game, gs *gamestates.GameStates) {
	g.Reset()
	g.Cards.TestingShuffle()
	g.DealBoard()
	gs.Reset()
	gs.SaveState(*g)
}

func ShowGameState(gs *gamestates.GameStates) {
	gs.PrintLast()
}

func ShowGameStates(gs *gamestates.GameStates) {
	gs.PrintAll()
}

func ChangeFlipCount(g *game.Game, gs *gamestates.GameStates) {
	g.SetFlipCount(1)
	fmt.Println("Easy mode.")
	gs.SaveState(*g)
}
