package gamestates

import (
	"solitaire/game"
	"solitaire/testutils"
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
	// Use an unshuffled deck to make sure the board is organized the same each time.
	// Made sure to test each type of move in game.
	// TestUndo uses these same moves.  If any are updated here, update below.
	gs := NewGameStates()
	game := createTestGameState()
	game.DealBoard()
	gs.SaveState(game)

	if len(gs.States) != 1 {
		t.Error("Expected 1 state after SaveState()")
	}
	testutils.AssertNoError(t, game.MoveFromBoardToStacks(0), "Move Board to Stacks failed")
	gs.SaveState(game)
	if len(gs.States) != 2 {
		t.Error("Expected 2 states after SaveState()")
	}
	if game.IsEqual(gs.States[0]) {
		t.Error("Saved state is not a deep copy of the original state")
	}
	testutils.AssertNoError(t, game.MoveFromBoardToStacks(2), "Move board to stack failed")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.MoveFromColumnToColumn(2, 4), "Move column to colum")
	gs.SaveState(game)
	if game.IsEqual(gs.States[2]) {
		t.Error("Saved state is not a deep copy of the original state")
	}
	testutils.AssertNoError(t, game.MoveFromColumnToColumn(5, 0), "")
	gs.SaveState(game)
	if game.IsEqual(gs.States[3]) {
		t.Error("Saved state is not a deep copy of the original state")
	}
	testutils.AssertNoError(t, game.NextDeckCard(), "")
	gs.SaveState(game)
	if game.IsEqual(gs.States[4]) {
		t.Error("Saved state is not a deep copy of the original state")
	}
	testutils.AssertNoError(t, game.NextDeckCard(), "")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.NextDeckCard(), "")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.MoveFromDeckToBoard(0), "Move deck to board")
	gs.SaveState(game)
	if game.IsEqual(gs.States[7]) {
		t.Error("Saved state is not a deep copy of the original state")
	}
	testutils.AssertNoError(t, game.MoveFromColumnToColumn(5, 0), "")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.MoveFromColumnToColumn(5, 4), "")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.NextDeckCard(), "Next card")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.MoveFromDeckToBoard(2), "Move deck to board")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.MoveFromDeckToStacks(), "Move deck to stack")
	gs.SaveState(game)
	if game.IsEqual(gs.States[12]) {
		t.Error("Saved state is not a deep copy of the original state")
	}
	testutils.AssertNoError(t, game.MoveFromColumnToColumn(2, 5), "Move column to column")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.MoveFromDeckToBoard(2), "Move deck to board")
	gs.SaveState(game)
	testutils.AssertNoError(t, game.SetFlipCount(1), "Set flip count to 1")
	gs.SaveState(game)
	// TODO: Finish testing SetFlipCount.
}

func TestUndo(t *testing.T) {
	// This is using the same moves as TestSaveState.  If any is updated there, update here.
	gs := NewGameStates()
	game := createTestGameState()
	game.DealBoard()
	gs.SaveState(game)

	game.MoveFromBoardToStacks(0)
	gs.SaveState(game)
	game.MoveFromBoardToStacks(2)
	gs.SaveState(game)
	game.MoveFromColumnToColumn(2, 4)
	gs.SaveState(game)
	game.MoveFromColumnToColumn(5, 0)
	gs.SaveState(game)
	game.NextDeckCard()
	gs.SaveState(game)
	game.NextDeckCard()
	gs.SaveState(game)
	game.NextDeckCard()
	gs.SaveState(game)
	game.MoveFromDeckToBoard(0)
	gs.SaveState(game)
	game.MoveFromColumnToColumn(5, 0)
	gs.SaveState(game)
	game.MoveFromColumnToColumn(5, 4)
	gs.SaveState(game)
	game.NextDeckCard()
	gs.SaveState(game)
	game.MoveFromDeckToBoard(2)
	gs.SaveState(game)
	game.MoveFromDeckToStacks()
	gs.SaveState(game)
	game.MoveFromColumnToColumn(2, 5)
	prevGame := game.DeepCopy()
	gs.SaveState(game)
	game.MoveFromDeckToBoard(2)
	gs.SaveState(game)

	stateSize := len(gs.States)
	undoState := gs.Undo()
	prevGame.Print()
	undoState.Print()
	if len(gs.States) != stateSize-1 {
		t.Error("State size should have been reduced by 1 after undo.")
	}
	if !undoState.IsEqual(prevGame) {
		t.Error("Undo did not move to the previous state.")
	}

	// Consider more tests for edge cases around multiple
	// undos and undos after resets.
}
