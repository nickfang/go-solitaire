package main

import (
	"fmt"
	"strconv"

	"solitaire/game"
	"solitaire/game/gamestates"

	"golang.org/x/exp/slices"
)

var validColumns = []string{"0", "1", "2", "3", "4", "5", "6"}

func NextCard(g *game.Game, gs gamestates.GameStates) {
	nextErr := g.NextDeckCard()
	if nextErr != nil {
		fmt.Println(nextErr.Error())
		return
	}
	DisplayGame(*g)
	gs.SaveState(*g)
}

func ResetGame(g *game.Game, gs gamestates.GameStates) {
	g.Reset()
	g.Cards.RandomShuffle()
	g.DealBoard()
	DisplayGame(*g)
	gs.Reset()
	gs.SaveState(*g)
}

func ShowHints(g *game.Game) {
	moves := g.GetDeckHints()
	moves = append(moves, g.GetStackHints()...)
	moves = append(moves, g.GetBoardHints()...)
	fmt.Println("Moves:", moves)
}

func Undo(g *game.Game, gs gamestates.GameStates) {
	if len(gs.States) <= 1 {
		fmt.Println("No moves to undo.")
	} else {
		lastGameState := gs.Undo()
		g.SetState(lastGameState)
	}
	DisplayGame(*g)
}

func ShowHelp() {
	fmt.Println("Commands: ")
	fmt.Println("  n - next card")
	fmt.Println("  d# - move from deck to column number")
	fmt.Println("  ds - move from deck to stacks")
	fmt.Println("  ## - move fromcolumn to column")
	fmt.Println("  r - reset")
	fmt.Println("  h - hints")
	fmt.Println("  fc1 - set flip count to 1 (easy mode)")
	fmt.Println("  u - undo")
	fmt.Println("  q - quit")
}

func MoveDeckToBoard(input1 string, g *game.Game, gs gamestates.GameStates) {
	if slices.Contains(validColumns, input1) {
		columnIndex, _ := strconv.ParseInt(input1, 10, 32)
		g.MoveFromDeckToBoard(int(columnIndex))
		DisplayGame(*g)
		gs.SaveState(*g)
	}
}

func MoveDeckToStacks(g *game.Game, gs gamestates.GameStates) {
	g.MoveFromDeckToStacks()
	DisplayGame(*g)
	gs.SaveState(*g)
}

func MoveBoardToStacks(input0 string, g *game.Game, gs gamestates.GameStates) {
	columnIndex, _ := strconv.ParseInt(input0, 10, 32)
	g.MoveFromBoardToStacks(int(columnIndex))
	DisplayGame(*g)
	gs.SaveState(*g)
}

func MoveColumnToColumn(input0, input1 string, g *game.Game, gs gamestates.GameStates) {
	if (slices.Contains(validColumns, input0) && slices.Contains(validColumns, input1)) && input0 != input1 {
		fromColumn, _ := strconv.ParseInt(input0, 10, 32)
		toColumn, _ := strconv.ParseInt(input1, 10, 32)
		g.MoveFromColumnToColumn(int(fromColumn), int(toColumn))
		DisplayGame(*g)
		gs.SaveState(*g)
	}
}

func DealTest(g *game.Game, gs gamestates.GameStates) {
	g.Reset()
	g.Cards.TestingShuffle()
	g.DealBoard()
	DisplayGame(*g)
	gs.Reset()
	gs.SaveState(*g)
}

func ShowGameStates(gs gamestates.GameStates) {
	gs.Print()
}

func ChangeFlipCount(g *game.Game, gs gamestates.GameStates) {
	g.SetFlipCount(1)
	fmt.Println("Easy mode.")
	DisplayGame(*g)
	gs.SaveState(*g)
}
