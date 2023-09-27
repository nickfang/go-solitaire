package main

import (
	"fmt"
	"strings"
)

func main() {
	game := newGame()
	game.cards.randomShuffle()
	game.cards, game.board, game.currentCardIndex = game.dealBoard()
	game.displayBoard()

	var input string
	game.displayCards()
	for {
		fmt.Scanln(&input)
		if strings.ToLower(input) == "q" {
			return
		}
		if strings.ToLower(input) == "n" {
			game = game.getNextCard()
			game.displayCards()
			continue
		}
		fmt.Println("Invalid Input.")
	}
}
