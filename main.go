package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	game := newGame()
	// game.cards.randomShuffle()
	game.cards, game.board, game.currentCardIndex = game.dealBoard()
	// game.setDebug(true)
	game.display()

	var i string
	fmt.Println(game.getDeckMoves())
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
		if input[:1] == "d" {
			to := string(input[1])
			if to == "s" {
				game.moveDeckToStacks()
			} else {
				columnIndex, _ := strconv.ParseInt(to, 10, 64)
				game.moveCurrentCard(int(columnIndex))
				game.display()
				fmt.Println(game.getDeckMoves())
			}
			continue
		}
		if input[:1] == "s" {
			// move from stacks to board.
		}
		fmt.Println("Invalid Input.", input)
	}
}
