package gamemanager

import (
	"errors"
	"sync"

	"solitaire/game"
	"solitaire/game/gamestates"

	"github.com/rs/xid"
)

type GameManager struct {
	games      map[string]*game.Game
	gameStates map[string]*gamestates.GameStates
	mu         sync.RWMutex
}

func NewGameManager() *GameManager {
	return &GameManager{
		games:      make(map[string]*game.Game),
		gameStates: make(map[string]*gamestates.GameStates),
		mu:         sync.RWMutex{},
	}
}

func (gm *GameManager) CreateGame() (string, error) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	gameId := xid.New().String()
	newGame := game.NewGame(gameId)
	newGameStates := gamestates.NewGameStates()
	newGame.Cards.RandomShuffle()
	newGameStates.SaveState(newGame)

	gm.games[gameId] = &newGame
	gm.gameStates[gameId] = &newGameStates

	return gameId, nil
}

func (gm *GameManager) GetGame(gameId string) (*game.Game, error) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()

	if game, ok := gm.games[gameId]; ok {
		return game, nil
	}
	return nil, errors.New("game not found")
}

func (gm *GameManager) GetGameStates(gameId string) (*gamestates.GameStates, error) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()

	if gameStates, ok := gm.gameStates[gameId]; ok {
		return gameStates, nil
	}
	return nil, errors.New("game states not found")
}
