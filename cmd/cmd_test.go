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
	sessionId, error := gm.CreateSession()
	if error != nil {
		t.Errorf("Error creating game: %s", error)
	}
	session, error := gm.GetSession(sessionId)
	if error != nil {
		t.Errorf("Error getting session: %s", error)
	}
	g := session.Game
	g.Cards.TestingShuffle()
	g.DealBoard()
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

	responseChan := make(chan gamemanager.GameResponse)
	for _, move := range moves {
		gr := gamemanager.GameRequest{SessionId: sessionId, Action: move, Response: responseChan}
		gm.Requests <- gr
		response := <-responseChan
		DisplayGame(*response.Game)
		error := response.Error
		if error != nil {
			t.Errorf("Error making move: %s - %s", move, error)
			return
		}
	}
	if !g.IsFinished() {
		t.Errorf("Game not won")
	}
	t.Fail()
}
