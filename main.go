package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
	"solitaire/game"
	"solitaire/gamestates"
)

func main() {
	game := game.NewGame()
	gameStates := gamestates.NewGameStates()
	game.Cards.RandomShuffle()
	game.Cards, game.Board, game.CurrentCardIndex = game.DealBoard()
	// game.SetDebug(true)
	gameStates.SaveState(game)
	game.Display()

	var i string
	validColumns := []string{"0", "1", "2", "3", "4", "5", "6"}
	for {
		fmt.Scanln(&i)
		input := strings.ToLower(i)
		if len(input) == 1 {
			if input == "q" {
				break
			}
			if input == "n" {
				game.NextDeckCard()
				game.Display()
				gameStates.SaveState(game)
				// fmt.Println(game.GetDeckMoves())
				continue
			}
			if input == "r" {
				game.Reset()
				game.Cards.RandomShuffle()
				game.Cards, game.Board, game.CurrentCardIndex = game.DealBoard()
				game.Display()
				gameStates.Reset()
				gameStates.SaveState(game)
				continue
			}
			if input == "h" {
				game.DisplayHints()
				continue
			}
			if input == "u" {
				if len(gameStates.States) == 0 {
					fmt.Printf("No moves to undo.")
				} else {
					lastGameState := gameStates.Undo()
					fmt.Print(lastGameState)
					game.SetState(lastGameState)
				}
				fmt.Printf("display new game state.")
				game.Display()
				continue
			}
		}
		if len(input) == 2 {
			input0 := string(input[0])
			input1 := string(input[1])
			if input0 == "d" {
				// move from deck to board
				to := input1
				if to == "s" {
					game.MoveFromDeckToStacks()
					game.Display()
					gameStates.SaveState(game)
				} else if slices.Contains(validColumns, to) {
					columnIndex, _ := strconv.ParseInt(to, 10, 32)
					game.MoveFromDeckToBoard(int(columnIndex))
					game.Display()
					gameStates.SaveState(game)
				} else {
					fmt.Println("Invalid Input.")
				}
				continue
			}
			if input[1] == 's' {
				// move from board to stacks
				from := input0
				if slices.Contains(validColumns, from) {
					columnIndex, _ := strconv.ParseInt(from, 10, 32)
					game.MoveFromBoardToStacks(int(columnIndex))
					game.Display()
					gameStates.SaveState(game)
					continue
				}
				fmt.Println("Invalid Input.")
			}
			if (slices.Contains(validColumns, input0) && slices.Contains(validColumns, input1)) && input0 != input1 {
				// move from column to column
				fromColumn, _ := strconv.ParseInt(input0, 10, 32)
				toColumn, _ := strconv.ParseInt(input1, 10, 32)
				game.MoveFromColumnToColumn(int(fromColumn), int(toColumn))
				game.Display()
				gameStates.SaveState(game)
				continue
			}
			if input[:1] == "s" {
				// move from stacks to board.
			}
		}
		fmt.Println("Invalid Input.", input)
	}
}
