package gamemanager

import (
	"errors"
	"solitaire/game"
	"solitaire/gamestates"

	"github.com/rs/xid"
)

type SessionRequest struct {
	Id     string
	Action string
}

type SessionResponse struct {
	Id      string
	Message string
	Error   error
}

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

func (gm *GameManager) SessionEngine() {
	for {
		sessionReq := <-gm.SessionReq
		switch sessionReq.Action {
		case "create":
			sessionId, _ := gm.CreateSession(WithRandomShuffle())
			session, error := gm.GetSession(sessionId)
			if error != nil || session == nil {
				gm.SessionRes <- SessionResponse{Id: sessionId, Message: "Session not created", Error: error}
				continue
			}
			gm.SessionRes <- SessionResponse{Id: sessionId, Message: "Session Created", Error: nil}
			continue
		case "create:test":
			sessionId, _ := gm.CreateSession(WithTestingShuffle())
			session, error := gm.GetSession(sessionId)
			if error != nil || session == nil {
				gm.SessionRes <- SessionResponse{Id: sessionId, Message: "Test session not created", Error: error}
				continue
			}
			gm.SessionRes <- SessionResponse{Id: sessionId, Message: "Test session Created", Error: nil}
			continue
		case "create:no-shuffle":
			sessionId, _ := gm.CreateSession()
			session, error := gm.GetSession(sessionId)
			if error != nil || session == nil {
				gm.SessionRes <- SessionResponse{Id: sessionId, Message: "Unshuffled session not created", Error: error}
				continue
			}
			gm.SessionRes <- SessionResponse{Id: sessionId, Message: "Unshuffled session Created", Error: nil}
			continue

			// I don't think we need to do a get.  Clients will communicate with the game through the game channels
		// case "get":
		// 	gm.GetSession(sessionReq.Id)
		// 	continue
		case "delete":
			error := gm.DeleteSession(sessionReq.Id)
			if error != nil {
				gm.SessionRes <- SessionResponse{Id: sessionReq.Id, Message: "Session not deleted", Error: error}
				continue
			}
			gm.SessionRes <- SessionResponse{Id: sessionReq.Id, Message: "Session Deleted", Error: error}
			continue
		case "quit":
			gm.SessionRes <- SessionResponse{Id: sessionReq.Id, Message: "Quitting", Error: nil}
			return
		}
	}
}
