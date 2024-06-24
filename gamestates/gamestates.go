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

func (s *GameStates) Reset() {
	s.States = s.States[:0]
}

func (s *GameStates) SaveState(state game.Game) {
	newState := state.DeepCopy()
	s.push(newState)
}

func (s *GameStates) Undo() game.Game {
	numStates := len(s.States)
	if numStates <= 1 {
		return s.States[0]
	}
	s.pop()
	lastIndex := len(s.States) - 1
	return s.States[lastIndex]
}

// for debugging
func (s *GameStates) PrintLast() {
	lastIndex := len(s.States) - 1
	s.States[lastIndex].Print()
}

func (s *GameStates) PrintAll() {
	for i, state := range s.States {
		state.Print()
		if i < len(s.States)-1 {
			println()
		}
	}
}
