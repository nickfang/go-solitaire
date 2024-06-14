package main

import (
	"fmt"
	"strings"

	"solitaire/game"
	"solitaire/game/gamestates"

	"golang.org/x/exp/slices"
)

var ValidColumns = []string{"1", "2", "3", "4", "5", "6", "7"}

func HandleMoves(game *game.Game, gameStates *gamestates.GameStates) {
	var i string
	// validColumns := []string{"1", "2", "3", "4", "5", "6", "7"}
	for {
		fmt.Scanln(&i)
		input := strings.ToLower(i)
		if input == "q" {
			break
		}
		if input == "n" {
			NextCard(game, gameStates)
			continue
		}
		if input == "r" {
			ResetGame(game, gameStates)
			continue
		}
		if input == "h" {
			DisplayHints(*game)
			continue
		}
		if input == "u" {
			Undo(game, gameStates)
			continue
		}
		if input == "?" {
			ShowHelp()
			continue
		}
		if input == "ds" {
			MoveDeckToStacks(game, gameStates)
			continue
		}
		if input == "rt" {
			DealTest(game, gameStates)
			continue
		}
		if input == "ss" {
			ShowGameStates(gameStates)
			continue
		}

		if input == "fc1" {
			ChangeFlipCount(game, gameStates)
			continue
		}
		// if moving to and/or from a column
		if len(input) == 2 {
			from := string(input[0])
			to := string(input[1])
			if from == "d" {
				if slices.Contains(ValidColumns, to) {
					MoveDeckToBoard(to, game, gameStates)
					continue
				}
			}
			if to == "s" {
				if slices.Contains(ValidColumns, from) {
					MoveBoardToStacks(from, game, gameStates)
					continue
				}
			}
			if (slices.Contains(ValidColumns, from) && slices.Contains(ValidColumns, to)) && from != to {
				MoveColumnToColumn(from, to, game, gameStates)
				continue
			}
			if from == "s" {
				// move from stacks to board.
				fmt.Printf("Not Implemented.\n")
			}
		}
		fmt.Println("Invalid Input.", input)
	}
}

func main() {
	game := game.NewGame()
	game.SetDebug(false)
	gameStates := gamestates.NewGameStates()
	game.Cards.RandomShuffle()
	game.DealBoard()
	gameStates.SaveState(game)
	DisplayGame(game)
	HandleMoves(&game, &gameStates)
}
