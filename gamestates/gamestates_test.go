package gamestates

import (
	"solitaire/game"
	"testing"
)

type Game game.Game

// Helper function to create a simple test game state
func createTestGameState() game.Game {
    return game.NewGame()
}

// func (g1 Game) compareTo(g2 game.Game) bool {
// 	if (g1.Cards == nil || g1.Board == nil || g1.Stacks == nil) {
// 		fmt.Printf("State cannot have a nil value %v, %v, %v", g1.Cards, g1.Board, g1.Stacks)
// 		return false
// 	}
// 	if (g2.Cards == nil || g2.Board == nil || g2.Stacks == nil) {
// 		fmt.Printf("State cannot have a nil value %v, %v, %v", g2.Cards, g2.Board, g2.Stacks)
// 		return false
// 	}
// 	if (len(g1.Cards) != len(g2.Cards)) {
// 		fmt.Printf("Cards do not match")
// 	}
// }

func TestNewGameStates(t *testing.T) {
	gs := NewGameStates()
	if len(gs.States) != 0 {
			t.Errorf("Expected empty GameStates, got %v", gs.States)
	}
}

func TestPushPop(t *testing.T) {
    gs := NewGameStates()
    state1 := createTestGameState()
    state2 := createTestGameState()

    gs.push(state1)
    gs.push(state2)

    poppedState := gs.pop()
    if !poppedState.IsEqual(state2) {
        t.Error("Popped state didn't match last pushed state")
    }

    poppedState = gs.pop()
    if !poppedState.IsEqual(state1) {
        t.Error("Popped state didn't match expected order")
    }
}

func TestReset(t *testing.T) {
    gs := NewGameStates()
    gs.push(createTestGameState())
    gs.Reset()

    if len(gs.States) != 0 {
        t.Error("GameStates should be empty after Reset()")
    }
}

func TestSaveState(t *testing.T) {
	gs := NewGameStates()
	game := createTestGameState()
    game.DealBoard()
	gs.SaveState(game)

	if len(gs.States) != 1 {
		t.Error("Expected 1 state after SaveState()")
	}
    game.MoveFromBoardToStacks(0)
    gs.SaveState(game)
    // Check for deep copy by comparing individual fields
    if len(gs.States) != 2 {
        t.Error("Expected 2 states after Save STate()")
    }
	if game.IsEqual(gs.States[0]) {
        t.Error("Saved state is not a deep copy of the original state")
    }
    game.MoveFromBoardToStacks(2)
    gs.SaveState(game)
    // Check for deep copy by comparing individual fields
	if game.IsEqual(gs.States[1]) {
        t.Error("Saved state is not a deep copy of the original state")
    }
    game.MoveFromColumnToColumn(2,4)
    gs.SaveState(game)
    if game.IsEqual(gs.States[2]) {
        t.Error("Saved state is not a deep copy of the original state")
    }
    game.MoveFromColumnToColumn(5,0)
    gs.SaveState(game)
    if game.IsEqual(gs.States[3]) {
        t.Error("Saved state is not a deep copy of the original state")
    }
    game.NextDeckCard()
    gs.SaveState(game)
    if game.IsEqual(gs.States[5]) {
        t.Error("Saved state is not a deep copy of the original state")
    }
    // game.NextDeckCard()
    // gs.SaveState(game)
    // game.NextDeckCard()
    // gs.SaveState(game)
    // game.NextDeckCard()
    // gs.SaveState(game)
    // game.MoveFromDeckToBoard(0)
    // gs.SaveState(game)
    // if game.Cards.IsEqual(gs.States[7].Cards) &&
    //     game.Board.IsEqual(gs.States[7].Board) &&
    //     game.Stacks.IsEqual(gs.States[7].Stacks) &&
    //     game.CurrentCardIndex == gs.States[1].CurrentCardIndex {
    //     t.Error("Saved state is not a deep copy of the original state")
    // }
    // game.NextDeckCard()
    // gs.SaveState(game)
    // game.MoveFromDeckToBoard(2)
    // gs.SaveState(game)
    // game.MoveFromDeckToStacks()
    // gs.SaveState(game)
    // if game.Cards.IsEqual(gs.States[10].Cards) &&
    //     game.Board.IsEqual(gs.States[10].Board) &&
    //     game.Stacks.IsEqual(gs.States[10].Stacks) &&
    //     game.CurrentCardIndex == gs.States[1].CurrentCardIndex {
    //     t.Error("Saved state is not a deep copy of the original state")
    // }
}

func TestUndo(t *testing.T) {
    gs := NewGameStates()
    state1 := createTestGameState()
    state2 := createTestGameState()
    // Modify state2 slightly to differentiate it

    gs.SaveState(state1)
    gs.SaveState(state2)

    // undoneState := gs.Undo()
    // if undoneState != state1 {
    //     t.Error("Undo didn't return the expected state")
    // }

    // Consider more tests for edge cases around multiple
    // undos and undos after resets.
}
