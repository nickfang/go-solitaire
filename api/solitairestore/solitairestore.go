package solitairestore

import (
	"solitaire/game"
	"solitaire/gamestates"
)

type SolitaireStore interface {
	SaveGame(currentGameState game.Game, previousGameStates gamestates.GameStates) (string, error)
	LoadGame(gameId string) (game.Game, gamestates.GameStates, error)
}

type GameStateData struct {
	CurrentGameState   game.Game             `json:"currentGameState"`
	PreviousGameStates gamestates.GameStates `json:"previousGameStates"`
}

func New() SolitaireStore {
	return &fileStore{}
}
