package main

// import "fmt"

func main() {
	game := newGame()
	game.cards.randomShuffle()
	game.cards, game.board, game.currentCardIndex = game.dealBoard()
	game.displayBoard()
	game.displayCards()

	for i := 0; i < 10; i++ {
		game = game.getNextCard()
		game.displayCards()
	}
}
