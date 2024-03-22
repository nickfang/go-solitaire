package gamestates

import (
	"solitaire/board"
	"solitaire/deck"
	"solitaire/game"
	"solitaire/stacks"
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
	newState := game.Game{}
	// Deep Copy Board (assuming board.Board is [][]deck.Card)
	newState.Board = make(board.Board, len(state.Board))
	for i, row := range state.Board {
			newState.Board[i] = make([]deck.Card, len(row))
			copy(newState.Board[i], row)
	}

	// Deep Copy Cards
	newState.Cards = make(deck.Cards, len(state.Cards))
	copy(newState.Cards, state.Cards)

	// Deep Copy Stacks (assuming stacks.Stacks is [][]deck.Card)
	newState.Stacks = make(stacks.Stacks, len(state.Stacks))
	for i, suitStack := range state.Stacks {
			newState.Stacks[i] = make([]deck.Card, len(suitStack))
			copy(newState.Stacks[i], suitStack)
	}
	newState.CurrentCardIndex = state.CurrentCardIndex
	s.push(newState)
}

func (s *GameStates) Undo() game.Game {
	numStates := len(s.States)
	if numStates == 0 {

	}
	if (numStates <= 1) {
		return s.States[0]
	}
	s.pop()
	lastIndex := len(s.States) - 1
	return s.States[lastIndex]
}