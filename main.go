package main

// import "fmt"

func main() {
	game := newGame()
	// for i := 0; i < len(deck.cards); i++ {
	// 	fmt.Println(deck.cards[i])
	// }
	game.randomShuffle()
	game.randomShuffle()
	game.printCards()
}
