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
   originalState := createTestGameState()
   gs.SaveState(originalState)

   if len(gs.States) != 1 {
       t.Error("Expected 1 state after SaveState()")
   }

   // (Add more detailed checks that the copied state is
   //  truly a deep copy by comparing individual fields)
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


func TestUndoJustUndid(t *testing.T) {
   gs := NewGameStates()
   state1 := createTestGameState()
   gs.SaveState(state1)
   gs.pop()
    // undoneState := gs.Undo()
    // undoneState2 := gs.Undo()
    // if undoneState2 != undoneState {
    //     t.Error("Undo didn't return the same state without a new push")
    // }

}
