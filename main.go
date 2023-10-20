package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
	"solitaire/game"
)

func main() {
	game := game.NewGame()
	// game.Cards.RandomShuffle()
	game.Cards, game.Board, game.CurrentCardIndex = game.DealBoard()
	// game.SetDebug(true)
	game.Display()

	var i string
	validColumns := []string{"0", "1", "2", "3", "4", "5", "6"}
	for {
		fmt.Scanln(&i)
		input := strings.ToLower(i)
		// input0 := string(input[0])
		// input1 := string(input[1])
		if len(input) == 1 {
			if input == "q" {
				break
			}
			if input == "n" {
				game.NextDeckCard()
				game.DisplayCards()
				fmt.Println(game.GetDeckMoves())
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
				} else if slices.Contains(validColumns, to) {
					columnIndex, _ := strconv.ParseInt(to, 10, 32)
					game.MoveFromDeckToBoard(int(columnIndex))
					game.Display()
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
				continue
			}
			if input[:1] == "s" {
				// move from stacks to board.
			}
		}
		fmt.Println("Invalid Input.", input)
	}
}
