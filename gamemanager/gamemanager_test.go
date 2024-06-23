package gamemanager

import (
	"testing"
)

func TestNewGameManager(t *testing.T) {
	gm := NewGameManager()
	if gm == nil {
		t.Error("Expected GameManager, got nil")
	}
	if len(gm.Sessions) != 0 {
		t.Error("Expected map[string]GameSession to be an empty string.")
	}
}

func TestCreateSession(t *testing.T) {
	gm := NewGameManager()
	gameId, error := gm.CreateSession()
	if error != nil {
		t.Errorf("Error creating game: %s", error)
	}
	if gameId == "" {
		t.Error("Expected gameId, got empty string")
	}
	if gm.Sessions[gameId].Game == nil {
		t.Error("Expected game, got nil")
	}
	if gm.Sessions[gameId].GameStates == nil {
		t.Error("Expected gameStates, got nil")
	}

}
