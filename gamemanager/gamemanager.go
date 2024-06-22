package gamemanager

import (
	"errors"
	"sync"

	"solitaire/game"
	"solitaire/gamestates"

	"github.com/rs/xid"
)

type GameSession struct {
	gameId     string
	game       *game.Game
	gameStates *gamestates.GameStates
}

type GameManager struct {
	Sessions map[string]GameSession
	mu       sync.RWMutex
}

func NewGameManager() *GameManager {
	return &GameManager{
		Sessions: make(map[string]GameSession),
		mu:       sync.RWMutex{},
	}
}

func (gm *GameManager) CreateGame() (string, error) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	gameId := xid.New().String()
	newGame := game.NewGame(gameId)
	newGame.Cards.RandomShuffle()
	newGameStates := gamestates.NewGameStates()
	newGameStates.SaveState(newGame)

	session := GameSession{gameId, &newGame, &newGameStates}
	gm.Sessions[gameId] = session

	return gameId, nil
}

func (gm *GameManager) GetSession(gameId string) (*GameSession, error) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()

	if session, ok := gm.Sessions[gameId]; ok {
		return &session, nil
	}
	return nil, errors.New("session not found")
}
