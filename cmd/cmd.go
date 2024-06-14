package main

import (
	"fmt"
	"strings"

	"solitaire/game"
	"solitaire/game/gamestates"

	"golang.org/x/exp/slices"
)

func HandleMoves(game *game.Game, gameStates gamestates.GameStates) {
	var i string
	validColumns := []string{"0", "1", "2", "3", "4", "5", "6"}
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
			ShowHints(game)
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
		if input == "ss" {
			ShowGameStates(gameStates)
			continue
		}
		if input == "rt" {
			fmt.Printf("Not Implemented.\n")
			// 	DealTest(game, gameStates)
			// 	continue
		}
		if input == "fc1" {
			ChangeFlipCount(game, gameStates)
			continue
		}
		if len(input) == 2 {
			from := string(input[0])
			to := string(input[1])
			if from == "d" {
				if to == "s" {
					MoveDeckToStacks(game, gameStates)
					continue
				} else if slices.Contains(validColumns, to) {
					MoveDeckToBoard(to, game, gameStates)
					continue
				}
			}
			if to == "s" {
				fmt.Println("from", from, "to", to)
				if slices.Contains(validColumns, from) {
					MoveBoardToStacks(from, game, gameStates)
					continue
				}
			}
			if (slices.Contains(validColumns, from) && slices.Contains(validColumns, to)) && from != to {
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
	HandleMoves(&game, gameStates)
}
