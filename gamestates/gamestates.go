package gamestates

import (
	"solitaire/game"
)

type GameStates struct {
	States []game.Game
	justUnDid bool
}

func NewGameStates() GameStates {
	return GameStates{}
}

func (s *GameStates) push(state game.Game) {
	s.States = append(s.States, state)
	s.justUnDid = false
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

func (s *GameStates) Reset() {
	s.States = s.States[:0]
	s.justUnDid = false
}

func (s *GameStates) SaveState(state game.Game) {
	newState := game.Game{}
	newState.Board = state.Board
	newState.Cards = state.Cards
	newState.Stacks = state.Stacks
	newState.CurrentCardIndex = state.CurrentCardIndex
	s.push(newState)
	s.justUnDid = false
}

func (s *GameStates) Undo() game.Game {
	if (len(s.States) == 1) {
		return s.States[0]
	}
	if !s.justUnDid {
		s.pop()
	}
	s.justUnDid = true
	return s.pop()
}