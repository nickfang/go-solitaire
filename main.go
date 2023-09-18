package main

// import "fmt"

func main() {
	game := newGame()

	game.cards.randomShuffle()
	game.dealBoard()
	game.board.print()
}
