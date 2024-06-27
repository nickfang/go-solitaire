package gamemanager

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"solitaire/game"
	"solitaire/gamestates"

	"github.com/rs/xid"
	"golang.org/x/exp/slices"
)

type GameRequest struct {
	SessionId string
	Action    string
	Data      interface{}
	Response  chan GameResponse
}

type GameResponse struct {
	Game    *game.Game
	Error   error
	Message string
}

type GameSession struct {
	Id         string
	Game       *game.Game
	GameStates *gamestates.GameStates
	Respons    chan GameResponse
}

type SessionEvent struct {
	SessionId   string
	EventType   string
	SessionData *GameSession
}

type GameManager struct {
	Sessions  map[string]*GameSession
	Requests  chan GameRequest
	Responses chan GameResponse
	Events    chan SessionEvent
	Mutex     sync.RWMutex
}

type GameRequestData struct {
}

type ClientOption func(*GameSession)

var ValidColumns = []string{"1", "2", "3", "4", "5", "6", "7"}

func NewGameManager() *GameManager {
	return &GameManager{
		Sessions:  make(map[string]*GameSession),
		Requests:  make(chan GameRequest),
		Responses: make(chan GameResponse),
		Events:    make(chan SessionEvent),
		Mutex:     sync.RWMutex{},
	}
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

// func (gm *GameManager) InitializeGame(sessionId string) error {
// 	gs, error := gm.GetSession(sessionId)
// 	if error != nil {
// 		return error
// 	}
// 	gs.Game.Cards.RandomShuffle()
// 	gs.Game.DealBoard()
// 	gs.GameStates.SaveState(*gs.Game)
// 	return nil
// }

// func (gm *GameManager) InitializeTestGame(sessionId string) error {
// 	gs, error := gm.GetSession(sessionId)
// 	if error != nil {
// 		return error
// 	}
// 	gs.Game.Cards.TestingShuffle()
// 	gs.Game.DealBoard()
// 	gs.GameStates.SaveState(*gs.Game)
// 	return nil
// }

func (gm *GameManager) GetSession(sessionId string) (*GameSession, error) {
	gm.Mutex.RLock()
	defer gm.Mutex.RUnlock()

	if session, ok := gm.Sessions[sessionId]; ok {
		return session, nil
	}
	return nil, errors.New("session not found")
}

func (gm *GameManager) DeleteSession(sessionId string) {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	delete(gm.Sessions, sessionId)
}

func (gm *GameManager) ProcessRequests() {
	for {
		req := <-gm.Requests
		if req.Action == "q" {
			req.Response <- GameResponse{Error: errors.New("quit")}
			gm.DeleteSession(req.SessionId)
			break
		}
		session, error := gm.GetSession(req.SessionId)
		if session == nil {
			req.Response <- GameResponse{Error: errors.New("session not found")}
			continue
		}
		if error != nil {
			req.Response <- GameResponse{Error: error}
			continue
		}
		error = HandleMoves(req.Action, session)
		if error != nil {
			req.Response <- GameResponse{Error: error}
			continue
		}
		req.Response <- GameResponse{Game: session.Game}
	}
	// close the channel
	close(gm.Requests)
}

func NextCard(g *game.Game, gs *gamestates.GameStates) error {
	nextErr := g.NextDeckCard()
	if nextErr != nil {
		return nextErr
	}
	gs.SaveState(*g)
	return nil
}

func GetHints(g *game.Game) []string {
	moves := g.GetDeckHints()
	moves = append(moves, g.GetStackHints()...)
	moves = append(moves, g.GetBoardHints()...)
	return moves
}

func Undo(g *game.Game, gs *gamestates.GameStates) error {
	if len(gs.States) <= 1 {
		return errors.New("no moves to undo")
	}
	lastGameState := gs.Undo()
	error := g.UpdateState(lastGameState)
	if error != nil {
		return error
	}
	return nil
}

func MoveDeckToBoard(input1 string, g *game.Game, gs *gamestates.GameStates) error {
	if !slices.Contains(ValidColumns, input1) {
		return errors.New("invalid column: " + input1)
	}
	columnIndex, _ := strconv.ParseInt(input1, 10, 32)
	error := g.MoveFromDeckToBoard(int(columnIndex - 1))
	if error != nil {
		return error
	}
	gs.SaveState(*g)
	return nil
}

func MoveDeckToStacks(g *game.Game, gs *gamestates.GameStates) error {
	error := g.MoveFromDeckToStacks()
	if error != nil {
		return error
	}
	gs.SaveState(*g)
	return nil
}

func MoveBoardToStacks(input0 string, g *game.Game, gs *gamestates.GameStates) error {
	columnIndex, _ := strconv.ParseInt(input0, 10, 32)
	error := g.MoveFromBoardToStacks(int(columnIndex - 1))
	if error != nil {
		return error
	}
	gs.SaveState(*g)
	return nil
}

func MoveColumnToColumn(input0, input1 string, g *game.Game, gs *gamestates.GameStates) error {
	if (!slices.Contains(ValidColumns, input0) || !slices.Contains(ValidColumns, input1)) || input0 == input1 {
		return errors.New("invalid move input")
	}

	fromColumn, _ := strconv.ParseInt(input0, 10, 32)
	toColumn, _ := strconv.ParseInt(input1, 10, 32)
	error := g.MoveFromColumnToColumn(int(fromColumn-1), int(toColumn-1))
	if error != nil {
		return error
	}

	gs.SaveState(*g)
	return nil
}

func DealTest(g *game.Game, gs *gamestates.GameStates) error {
	g.Reset()
	g.Cards.TestingShuffle()
	g.DealBoard()
	gs.Reset()
	gs.SaveState(*g)
	return nil
}

func ShowGameState(gs *gamestates.GameStates) {
	gs.PrintLast()
}

func ShowGameStates(gs *gamestates.GameStates) error {
	gs.PrintAll()
	return nil
}

func ChangeFlipCount(g *game.Game, gs *gamestates.GameStates) error {
	error := g.SetFlipCount(1)
	if error != nil {
		return error
	}
	fmt.Println("Easy mode.")
	gs.SaveState(*g)
	return nil
}

func HandleMoves(input string, session *GameSession) error {
	game := session.Game
	gameStates := session.GameStates
	if input == "n" {
		return NextCard(game, gameStates)
	}
	if input == "h" {
		GetHints(game)
		return nil
	}
	if input == "u" {
		return Undo(game, gameStates)
	}
	if input == "ds" {
		error := MoveDeckToStacks(game, gameStates)
		if error != nil {
			return error
		}
		if game.IsFinished() {
			return errors.New("finished")
		}
		return nil
	}
	if input == "rt" {
		return DealTest(game, gameStates)
	}
	if input == "ss" {
		return ShowGameStates(gameStates)
	}

	if input == "fc1" {
		return ChangeFlipCount(game, gameStates)
	}
	// if moving to and/or from a column
	if len(input) == 2 {
		from := string(input[0])
		to := string(input[1])
		if from == "d" {
			if slices.Contains(ValidColumns, to) {
				error := MoveDeckToBoard(to, game, gameStates)
				if error != nil {
					return error
				}
				return nil
			}
		}
		if to == "s" {
			if slices.Contains(ValidColumns, from) {
				error := MoveBoardToStacks(from, game, gameStates)
				if error != nil {
					return error
				}
				if game.IsFinished() {
					return errors.New("finished")
				}
				return nil
			}
		}
		if (slices.Contains(ValidColumns, from) && slices.Contains(ValidColumns, to)) && from != to {
			error := MoveColumnToColumn(from, to, game, gameStates)
			if error != nil {
				return error
			}
			return nil
		}
		if from == "s" {
			// move from stacks to board.
			// fmt.Printf("Not Implemented.\n")
			return fmt.Errorf("not Implemented: %s", input)
		}
	}
	return fmt.Errorf(`invalid Input: %s`, input)
}
