package main

import (
	"fmt"
	"strings"

	"solitaire/game"
	"solitaire/game/gamestates"

	"golang.org/x/exp/slices"
)

var ValidColumns = []string{"1", "2", "3", "4", "5", "6", "7"}

func HandleMoves(input string, game *game.Game, gameStates *gamestates.GameStates) error {
	if input == "q" {
		return fmt.Errorf("quitting")
	}
	if input == "n" {
		NextCard(game, gameStates)
		return nil
	}
	if input == "r" {
		ResetGame(game, gameStates)
		return nil
	}
	if input == "h" {
		DisplayHints(*game)
		return nil
	}
	if input == "u" {
		Undo(game, gameStates)
		return nil
	}
	if input == "?" {
		ShowHelp()
		return nil
	}
	if input == "ds" {
		MoveDeckToStacks(game, gameStates)
		return nil
	}
	if input == "rt" {
		DealTest(game, gameStates)
		return nil
	}
	if input == "ss" {
		ShowGameStates(gameStates)
		return nil
	}

	if input == "fc1" {
		ChangeFlipCount(game, gameStates)
		return nil
	}
	// if moving to and/or from a column
	if len(input) == 2 {
		from := string(input[0])
		to := string(input[1])
		if from == "d" {
			if slices.Contains(ValidColumns, to) {
				MoveDeckToBoard(to, game, gameStates)
				return nil
			}
		}
		if to == "s" {
			if slices.Contains(ValidColumns, from) {
				MoveBoardToStacks(from, game, gameStates)
				return nil
			}
		}
		if (slices.Contains(ValidColumns, from) && slices.Contains(ValidColumns, to)) && from != to {
			MoveColumnToColumn(from, to, game, gameStates)
			return nil
		}
		if from == "s" {
			// move from stacks to board.
			// fmt.Printf("Not Implemented.\n")
			return fmt.Errorf("not Implemented: %s", input)
		}
	}
	return fmt.Errorf(`invalid Input: %s`, input)
}

func main() {
	game := game.NewGame()
	game.SetDebug(false)
	gameStates := gamestates.NewGameStates()
	game.Cards.RandomShuffle()
	game.DealBoard()
	gameStates.SaveState(game)
	DisplayGame(game)
	var i string
	// validColumns := []string{"1", "2", "3", "4", "5", "6", "7"}
	for {
		fmt.Scanln(&i)
		input := strings.ToLower(i)
		err := HandleMoves(input, &game, &gameStates)
		if err != nil {
			if err.Error() == "quitting" {
				break
			}
			fmt.Println(err)
		}
	}
}
