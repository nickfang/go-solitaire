package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	game := newGame()
	// game.cards.randomShuffle()
	game.cards, game.board, game.currentCardIndex = game.dealBoard()
	// game.setDebug(true)
	game.display()

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
				game.nextDeckCard()
				game.displayCards()
				fmt.Println(game.getDeckMoves())
				continue
			}
		}
		if len(input) == 2 {
			input0 := string(input[0])
			input1 := string(input[1])
			if input0 == "d" {
				to := input1
				if to == "s" {
					game.moveFromDeckToStacks()
					game.display()
				} else if slices.Contains(validColumns, to) {
					columnIndex, _ := strconv.ParseInt(to, 10, 32)
					game.moveFromDeckToBoard(int(columnIndex))
					game.display()
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
					game.moveFromBoardToStacks(int(columnIndex))
					game.display()
					continue
				}
				fmt.Println("Invalid Input.")
			}
			if (slices.Contains(validColumns, input0) && slices.Contains(validColumns, input1)) && input0 != input1 {
				fromColumn, _ := strconv.ParseInt(input0, 10, 32)
				toColumn, _ := strconv.ParseInt(input1, 10, 32)
				game.moveFromBoardToBoard(int(fromColumn), int(toColumn))
				continue
			}
			if input[:1] == "s" {
				// move from stacks to board.
			}
		}
		fmt.Println("Invalid Input.", input)
	}
}
