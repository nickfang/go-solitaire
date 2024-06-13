package gamemanagerapi

import (
	"fmt"
	"strconv"
	"strings"

	"solitaire/game"
	"solitaire/game/deck"
	"solitaire/game/gamestates"

	"golang.org/x/exp/slices"
)

func CreateGame() (game.Game, gamestates.GameStates) {
	game := game.NewGame()
	gameStates := gamestates.NewGameStates()
	game.Cards.RandomShuffle()
	game.DealBoard()
	gameStates.SaveState(game)
	return game, gameStates
}

func HandleMoves(game *game.Game, gameStates gamestates.GameStates) {
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
				nextErr := game.NextDeckCard()
				if nextErr != nil {
					fmt.Println(nextErr.Error())
					continue
				}
				game.Display()
				gameStates.SaveState(*game)
				continue
			}
			if input == "r" {
				game.Reset()
				game.Cards.RandomShuffle()
				game.DealBoard()
				game.Display()
				gameStates.Reset()
				gameStates.SaveState(*game)
				continue
			}
			if input == "h" {
				game.DisplayHints()
				continue
			}
			if input == "u" {
				if len(gameStates.States) <= 1 {
					fmt.Println("No moves to undo.")
				} else {
					lastGameState := gameStates.Undo()
					game.SetState(lastGameState)
				}
				game.Display()
				continue
			}
			if input == "?" {
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
					gameStates.SaveState(*game)
				} else if slices.Contains(validColumns, to) {
					columnIndex, _ := strconv.ParseInt(to, 10, 32)
					game.MoveFromDeckToBoard(int(columnIndex))
					game.Display()
					gameStates.SaveState(*game)
				} else {
					fmt.Println("Invalid Input.")
				}
				continue
			}
			if input1 == "s" {
				// move from board to stacks
				from := input0
				columnIndex, _ := strconv.ParseInt(from, 10, 32)
				if slices.Contains(validColumns, from) {
					game.MoveFromBoardToStacks(int(columnIndex))
					game.Display()
					gameStates.SaveState(*game)
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
				gameStates.SaveState(*game)
				continue
			}
			if input == "rt" {
				game.Reset()
				deck, err := deck.NewTestingDeck()
				if err != nil {
					fmt.Println(err)
				}
				game.Cards = deck
				game.DealBoard()
				game.Display()
				gameStates.Reset()
				gameStates.SaveState(*game)
				continue
			}
			if input == "ss" {
				gameStates.Print()
				continue
			}
			if input0 == "s" {
				fmt.Printf("Not Implemented.\n")
				// move from stacks to board.
			}
		}
		if len(input) == 3 {
			if input == "fc1" {
				game.SetFlipCount(1)
				fmt.Println("Easy mode.")
				game.Display()
				gameStates.SaveState(*game)
				continue
			}
		}
		fmt.Println("Invalid Input.", input)
	}
}
