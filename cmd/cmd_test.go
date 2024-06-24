package main

import (
	"fmt"
	"solitaire/game"
	"solitaire/gamemanager"

	"testing"
)

func TestDealBoard(t *testing.T) {
	g := game.NewGame("")
	fmt.Println(g)
	g.DealBoard()
	// g.SetDebug(true)
	// g.board.Display()
}

func TestGetCardDisplay(t *testing.T) {
	cardDisplays := []string{}
	expected := []string{" A♠", " 2♠", " 3♠", " 4♠", " 5♠", " 6♠", " 7♠", " 8♠", " 9♠", "10♠", " J♠", " Q♠", " K♠", " A♥", " 2♥", " 3♥", " 4♥", " 5♥", " 6♥", " 7♥", " 8♥", " 9♥", "10♥", " J♥", " Q♥", " K♥", " A♣", " 2♣", " 3♣", " 4♣", " 5♣", " 6♣", " 7♣", " 8♣", " 9♣", "10♣", " J♣", " Q♣", " K♣", " A♦", " 2♦", " 3♦", " 4♦", " 5♦", " 6♦", " 7♦", " 8♦", " 9♦", "10♦", " J♦", " Q♦", " K♦"}

	for _, suit := range []string{"Spades", "Hearts", "Clubs", "Diamonds"} {
		for j := 1; j <= 13; j++ {
			cardDisplays = append(cardDisplays, getCardDisplay(j, suit))
		}
	}
	for i, cardDisplay := range cardDisplays {
		if cardDisplay != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], cardDisplay)
		}
	}
}

func TestFullGame(t *testing.T) {
	// This really only needs to test that the move strings call the correct functions.
	gm := gamemanager.NewGameManager()
	go gm.ProcessRequests()

	sessionId, error := gm.CreateSession()
	if error != nil {
		t.Errorf("Error creating game: %s", error)
	}

	error = gm.InitializeTestGame(sessionId)
	if error != nil {
		t.Errorf("Error initializing test game: %s", error)
	}

	session, error := gm.GetSession(sessionId)
	if error != nil {
		t.Errorf("Error getting session: %s", error)
	}
	g := *session.Game
	moves := []string{
		"ds", "ds", "ds", "n", "ds", "ds", "ds", "n",
		"ds", "ds", "ds", "n", "ds", "ds", "ds", "n",
		"ds", "ds", "ds", "n", "ds", "ds", "ds", "n",
		"ds", "ds", "ds", "n", "ds", "ds", "ds",
		"7s", "7s", "6s", "76", "64", "7s", "5s", "7s", "3s",
		"57", "53", "52", "21", "52", "43", "3s", "3s", "6s",
		"7s", "6s", "3s", "7s", "3s", "4s", "6s", "3s", "2s",
		"1s", "4s", "6s", "2s", "74", "12", "2s", "4s", "2s", "3s", "4s",
	}

	responseChan := make(chan gamemanager.GameResponse, 10)
	for _, move := range moves {
		gr := gamemanager.GameRequest{SessionId: sessionId, Action: move, Response: responseChan}
		gm.Requests <- gr

		response := <-responseChan
		error := response.Error
		if error != nil {
			t.Errorf("Error making move: %s - %s", move, error)
			return
		}
	}
	if !g.IsFinished() {
		t.Errorf("Game not won")
	}
}

func TestInvalidMoves(t *testing.T) {
	gm := gamemanager.NewGameManager()
	go gm.ProcessRequests()

	sessionId, error := gm.CreateSession()
	if error != nil {
		t.Errorf("Error creating game: %s", error)
	}

	error = gm.InitializeTestGame(sessionId)
	if error != nil {
		t.Errorf("Error initializing test game: %s", error)
	}

	responseChan := make(chan gamemanager.GameResponse)
	gr := gamemanager.GameRequest{SessionId: sessionId, Action: "ds", Response: responseChan}
	gm.Requests <- gr
	gm.Requests <- gr
	gm.Requests <- gr
	gm.Requests <- gr
	response := <-responseChan
	response = <-responseChan
	response = <-responseChan
	response = <-responseChan
	if response.Error.Error() != "no cards in the deck" {
		t.Errorf("Expected error: no cards in the deck")
	}
	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "12", Response: responseChan}
	gm.Requests <- gr
	response = <-responseChan
	if response.Error.Error() != "invalid board move" {
		t.Errorf("Expected error: invalid board move")
	}
	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "1s", Response: responseChan}
	gm.Requests <- gr
	response = <-responseChan
	if response.Error.Error() != "invalid move" {
		t.Errorf("Expected error: invalid move")
	}
	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "d1", Response: responseChan}
	gm.Requests <- gr
	response = <-responseChan
	if response.Error.Error() != "invalid move" {
		t.Errorf("Expected error: invalid move")
	}
	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "n", Response: responseChan}
	gm.Requests <- gr
	response = <-responseChan
	if response.Error != nil {
		t.Errorf("Expected no error")
	}

}
