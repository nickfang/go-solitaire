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

type GameSession struct {
	Id         string
	Game       *game.Game
	GameStates *gamestates.GameStates
}

type GameManager struct {
	Sessions map[string]GameSession
	mu       sync.RWMutex
	requests chan GameRequest
}

type GameResponse struct {
	Game    *game.Game
	Error   error
	Message string
}

type GameRequestData struct {
}

type GameRequest struct {
	SessionId string
	Action    string
	Data      interface{}
	Response  chan GameResponse
}

var ValidColumns = []string{"1", "2", "3", "4", "5", "6", "7"}

func NewGameManager() *GameManager {
	return &GameManager{
		Sessions: make(map[string]GameSession),
		mu:       sync.RWMutex{},
		requests: make(chan GameRequest),
	}
}

func (gm *GameManager) CreateSession() (string, error) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	sessionId := xid.New().String()
	game := game.NewGame(sessionId)
	error := game.Cards.RandomShuffle()
	if error != nil {
		return "", error
	}
	game.DealBoard()
	gameStates := gamestates.NewGameStates()
	gameStates.SaveState(game)
	newSession := GameSession{
		Id:         sessionId,
		Game:       &game,
		GameStates: &gameStates,
	}
	gm.Sessions[sessionId] = newSession
	return sessionId, nil
}

func (gm *GameManager) GetSession(sessionId string) (*GameSession, error) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()

	if session, ok := gm.Sessions[sessionId]; ok {
		return &session, nil
	}
	return nil, errors.New("session not found")
}

func (gm *GameManager) DeleteSession(sessionId string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	delete(gm.Sessions, sessionId)
}

func (ge *GameManager) ProcessRequests() {
	for {
		req := <-ge.requests
		if req.Action == "kill" {
			break
		}
		if req.Action == "q" {
			ge.DeleteSession(req.SessionId)
			continue
		}
		session, error := ge.GetSession(req.SessionId)
		if session == nil {
			req.Response <- GameResponse{Error: errors.New("session not found")}
			continue
		}
		if error != nil {
			req.Response <- GameResponse{Error: error}
			continue
		}
		error = HandleMoves(req.Action, *session)
		if error != nil {
			req.Response <- GameResponse{Error: error}
			continue
		}
	}
	// close the channel
	close(ge.requests)
}

func NextCard(g *game.Game, gs *gamestates.GameStates) error {
	nextErr := g.NextDeckCard()
	if nextErr != nil {
		return nextErr
	}
	gs.SaveState(*g)
	return nil
}

func ResetGame(g *game.Game, gs *gamestates.GameStates) error {
	error := g.Reset()
	if error != nil {
		return error
	}
	gs.Reset()
	gs.SaveState(*g)
	return nil
}

func GetHints(g *game.Game) {
	moves := g.GetDeckHints()
	moves = append(moves, g.GetStackHints()...)
	moves = append(moves, g.GetBoardHints()...)
	fmt.Println("Moves:", moves)
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

func DealTest(g *game.Game, gs *gamestates.GameStates) {
	g.Reset()
	g.Cards.TestingShuffle()
	g.DealBoard()
	gs.Reset()
	gs.SaveState(*g)
}

func ShowGameState(gs *gamestates.GameStates) {
	gs.PrintLast()
}

func ShowGameStates(gs *gamestates.GameStates) {
	gs.PrintAll()
}

func ChangeFlipCount(g *game.Game, gs *gamestates.GameStates) {
	g.SetFlipCount(1)
	fmt.Println("Easy mode.")
	gs.SaveState(*g)
}

func HandleMoves(input string, session GameSession) error {
	game := session.Game
	gameStates := session.GameStates
	if input == "n" {
		NextCard(game, gameStates)
		return nil
	}
	if input == "r" {
		ResetGame(game, gameStates)
		return nil
	}
	if input == "h" {
		GetHints(game)
		return nil
	}
	if input == "u" {
		err := Undo(game, gameStates)
		return err
	}
	if input == "ds" {
		MoveDeckToStacks(game, gameStates)
		return nil
	}
	if input == "rt" {
		DealTest(game, gameStates)
		return nil
	}
	if input == "ss" {
		ShowGameStates(gameStates)
		return nil
	}

	if input == "fc1" {
		ChangeFlipCount(game, gameStates)
		return nil
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
