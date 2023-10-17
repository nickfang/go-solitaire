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
	game.setDebug(true)
	game.display()

	var i string
	validColumns := []string{"0", "1", "2", "3", "4", "5", "6"}
	for {
		fmt.Scanln(&i)
		input := strings.ToLower(i)
		if input == "q" {
			break
		}
		if input == "n" {
			game.nextDeckCard()
			game.displayCards()
			fmt.Println(game.getDeckMoves())
			continue
		}
		if string(input[0]) == "d" {
			to := string(input[1])
			if to == "s" {
				game.moveFromDeckToStacks()
				game.display()
			} else if slices.Contains(validColumns, to) {
				columnIndex, _ := strconv.ParseInt(to, 10, 64)
				fmt.Println("columnIndex", columnIndex)
				game.moveFromDeckToBoard(int(columnIndex))
				game.display()
			} else {
				fmt.Println("Invalid Input.")
			}
			continue
		}
		if input[1] == 's' {
			// move from board to stacks
			from := string(input[0])
			fmt.Println(from, validColumns, slices.Contains(validColumns, from))
			if slices.Contains(validColumns, from) {
				columnIndex, _ := strconv.ParseInt(from, 10, 64)
				game.moveFromBoardToStacks(int(columnIndex))
				game.display()
				continue
			}
			fmt.Println("Invalid Input.")
		}
		if input[:1] == "s" {
			// move from stacks to board.
		}
		fmt.Println("Invalid Input.", input)
	}
}
