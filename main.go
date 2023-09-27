package main

import "fmt"

// import "fmt"

func main() {
	game := newGame()
	game.cards.randomShuffle()
	game.cards, game.board, game.currentCardIndex = game.dealBoard()
	game.displayBoard()

	for i := 0; i < 10; i++ {
		game.displayCards()
		fmt.Println(game.getDeckMoves())
		game = game.getNextCard()
	}
}
