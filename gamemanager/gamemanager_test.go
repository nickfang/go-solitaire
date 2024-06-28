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
