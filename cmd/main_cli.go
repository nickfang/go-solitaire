package main

import (
	"solitaire/game"
	"solitaire/game/gamestates"
	"solitaire/gamemanagerapi"
)

func main() {
	game := game.NewGame()
	game.SetDebug(false)
	gameStates := gamestates.NewGameStates()
	game.Cards.RandomShuffle()
	game.DealBoard()
	gameStates.SaveState(game)
	game.Display()
	gamemanagerapi.HandleMoves(&game, gameStates)
}
