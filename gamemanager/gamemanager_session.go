package gamemanager

import (
	"errors"
	"solitaire/game"
	"solitaire/gamestates"

	"github.com/rs/xid"
)

func (gm *GameManager) CreateSession(options ...ClientOption) (string, error) {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	sessionId := xid.New().String()
	game := game.NewGame(sessionId)
	gameStates := gamestates.NewGameStates()
	newSession := &GameSession{
		Id:         sessionId,
		Game:       &game,
		GameStates: &gameStates,
	}

	for _, option := range options {
		option(newSession)
	}

	newSession.Game.DealBoard()
	newSession.GameStates.SaveState(*newSession.Game)

	gm.Sessions[sessionId] = newSession
	return sessionId, nil
}

func WithRandomShuffle() ClientOption {
	return func(gs *GameSession) {
		gs.Game.Cards.RandomShuffle()
	}
}

func WithTestingShuffle() ClientOption {
	return func(gs *GameSession) {
		gs.Game.Cards.TestingShuffle()
	}
}

func (gm *GameManager) GetSession(sessionId string) (*GameSession, error) {
	gm.Mutex.RLock()
	defer gm.Mutex.RUnlock()

	if session, ok := gm.Sessions[sessionId]; ok {
		return session, nil
	}
	return nil, errors.New("session not found")
}

func (gm *GameManager) DeleteSession(sessionId string) error {
	if _, ok := gm.Sessions[sessionId]; !ok {
		return errors.New("session not found")
	}
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	delete(gm.Sessions, sessionId)
	return nil
}
