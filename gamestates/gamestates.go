package gamestates

import (
	"solitaire/game"
)

type GameStates struct {
	States []game.Game
}

func NewGameStates() GameStates {
	return GameStates{}
}

func (s *GameStates) push(state game.Game) {
	s.States = append(s.States, state)
}

func (s *GameStates) pop() game.Game {
	var defaultState game.Game
	if len(s.States) == 0 {
		return defaultState
	}
	lastIndex := len(s.States) - 1
	lastState := s.States[lastIndex]
	s.States = s.States[:lastIndex]
	return lastState

}

func (s *GameStates) SaveState(state game.Game) {
	s.push(state)
}

func (s *GameStates) Undo() game.Game {
	return s.pop()
}