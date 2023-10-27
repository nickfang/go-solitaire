package main

import "fmt"

func main() {
	game := newGame()

	game.cards.randomShuffle()
	game.cards.displayAll()
	game.cards, game.board = game.dealBoard()
	fmt.Println(game.board[5])
	// fmt.Println(game.board[1])
	// game.cards.display()
	game.board.display()
}
