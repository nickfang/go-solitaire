package gamemanager

import (
	"testing"
)

func TestCreateSession(t *testing.T) {
	gm := NewGameManager()
	// Test with no options.  Does not shuffle deck before dealing.
	sessionId, error := gm.CreateSession()
	if error != nil {
		t.Errorf("Error creating game: %s", error)
	}
	if sessionId == "" {
		t.Error("Expected sessionId, got empty string")
	}
	if gm.Sessions[sessionId].Game == nil {
		t.Error("Expected game, got nil")
	}
	if gm.Sessions[sessionId].GameStates == nil {
		t.Error("Expected gameStates, got nil")
	}

	sessionId, error = gm.CreateSession(WithRandomShuffle())
	if error != nil {
		t.Error("Expected no error.", error)
	}
	if sessionId == "" {
		t.Error("Expected sessionId, got empty string")
	}

	sessionId, error = gm.CreateSession(WithTestingShuffle())
	if error != nil {
		t.Error("Expected no error.", error)
	}
	if sessionId == "" {
		t.Error("Expected sessionId, got empty string")
	}

}

func TestGetSession(t *testing.T) {
	gm := NewGameManager()
	sessionId, _ := gm.CreateSession()
	session, error := gm.GetSession(sessionId)
	if error != nil {
		t.Errorf("Error getting session: %s", error)
	}
	if session == nil {
		t.Error("Expected session, got nil")
	}
	session, error = gm.GetSession("invalidId")
	if error.Error() != "session not found" {
		t.Error("Expected error, got nil", error)
	}
	if session != nil {
		t.Error("Expected nil session, got session")
	}
}

func TestDeleteSession(t *testing.T) {
	gm := NewGameManager()
	sessionId, _ := gm.CreateSession()
	error := gm.DeleteSession(sessionId)
	if error != nil {
		t.Errorf("Error deleting session: %s", error)
	}
	if gm.Sessions[sessionId] != nil {
		t.Error("Expected nil session, got session")
	}

	error = gm.DeleteSession("invalidId")
	if error.Error() != "session not found" {
		t.Error("Expected error, got nil", error)
	}

}
