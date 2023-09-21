package main

// import "fmt"

func main() {
	game := newGame()

	game.cards.randomShuffle()
	game.cards.displayAll()
	game.cards, game.board = game.dealBoard()
	// game.cards.display()
	game.board.display()
}
