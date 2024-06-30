package main

import (
	"fmt"
	"io"
	"os"
	"solitaire/gamemanager"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	oldArgs := os.Args
	oldStdin := os.Stdin
	oldStdout := os.Stdout

	defer func() {
		os.Args = oldArgs
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	os.Args = []string{"go run ."}
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stdin = r

	// run some commands, end with 'q' or test times out.
	fmt.Fprintln(w, "n")
	fmt.Fprintln(w, "h")
	fmt.Fprintln(w, "?")
	fmt.Fprintln(w, "r")
	fmt.Fprintln(w, "rt")
	fmt.Fprintln(w, "q")

	main()

	// must close write before getting read or test times out
	w.Close()

	out, _ := io.ReadAll(r)
	r.Close()

	// to print the output, redirect stdout to the original value
	// os.Stdout = oldStdout
	// fmt.Println(string(out))

	if len(out) == 0 {
		t.Errorf("Expected non-empty output, but got empty output")
	}

	lines := strings.Split(string(out), "\n")
	lastLine := lines[len(lines)-2]
	if lastLine != "Quitting..." {
		t.Errorf("Expected last line to be 'Quitting...', but got '%s'", lastLine)
	}
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

func TestGetCardDisplayInvalidCard(t *testing.T) {
	// Test with an invalid card value
	cardDisplay := getCardDisplay(14, "Spades")
	expected := "Invalid Value for Card"
	if cardDisplay != expected {
		t.Errorf("Expected %s, got %s", expected, cardDisplay)
	}

	// Test with an invalid suit
	cardDisplay = getCardDisplay(5, "InvalidSuit")
	expected = "Invalid Suit for Card"
	if cardDisplay != expected {
		t.Errorf("Expected %s, got %s", expected, cardDisplay)
	}
}

func TestFullGame(t *testing.T) {
	// This really only needs to test that the move strings call the correct functions.
	gm := gamemanager.NewGameManager()
	// defer gamemanager.CloseManager(gm)
	go gm.GameEngine()
	go gm.SessionEngine()

	gm.SessionReq <- gamemanager.SessionRequest{Action: "create:test"}
	sessionRes := <-gm.SessionRes
	if sessionRes.Error != nil {
		t.Errorf("Error creating session: %s", sessionRes.Error)
	} else {
		fmt.Println(sessionRes.Message)
	}
	sessionId := sessionRes.Id
	fmt.Println("Session ID: ", sessionId)

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

	for i, move := range moves {
		gr := gamemanager.GameRequest{SessionId: sessionId, Action: move}
		gm.GameReq <- gr

		response := <-gm.GameRes
		error := response.Error
		if i == len(moves)-1 {
			fmt.Println(i, len(moves)-1, error.Error(), "finished")
			if error.Error() != "finished" {
				t.Errorf("error should be: Congrat! You won!")
				return
			}
		} else {
			if error != nil {
				t.Errorf("Error making move: %s - %s", move, error)
				return
			}

		}
	}
	// if !g.IsFinished() {
	// 	t.Errorf("Game not won")
	// }
}

func TestInvalidMoves(t *testing.T) {
	gm := gamemanager.NewGameManager()
	// defer gamemanager.CloseManager(gm)
	go gm.GameEngine()
	go gm.SessionEngine()

	gm.SessionReq <- gamemanager.SessionRequest{Action: "create:test"}
	sessionRes := <-gm.SessionRes
	if sessionRes.Error != nil {
		t.Errorf("Error creating session: %s", sessionRes.Error)
	} else {
		fmt.Println(sessionRes.Message)
	}
	sessionId := sessionRes.Id

	move := "test"
	gr := gamemanager.GameRequest{SessionId: sessionId, Action: move}
	gm.GameReq <- gr
	response := <-gm.GameRes
	if response.Error == nil {
		t.Errorf("Expected error for invalid move, but got no error")
	}

	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "ds"}
	gm.GameReq <- gr
	response = <-gm.GameRes
	gm.GameReq <- gr
	response = <-gm.GameRes
	gm.GameReq <- gr
	response = <-gm.GameRes
	gm.GameReq <- gr
	response = <-gm.GameRes
	if response.Error.Error() != "no cards in the deck" {
		t.Errorf("Expected error: no cards in the deck")
	}
	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "12"}
	gm.GameReq <- gr
	response = <-gm.GameRes
	if response.Error.Error() != "invalid board move" {
		t.Errorf("Expected error: invalid board move")
	}
	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "1s"}
	gm.GameReq <- gr
	response = <-gm.GameRes
	if response.Error.Error() != "invalid move" {
		t.Errorf("Expected error: invalid move")
	}
	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "d1"}
	gm.GameReq <- gr
	response = <-gm.GameRes
	if response.Error.Error() != "invalid move" {
		t.Errorf("Expected error: invalid move")
	}
	gr = gamemanager.GameRequest{SessionId: sessionId, Action: "n"}
	gm.GameReq <- gr
	response = <-gm.GameRes
	if response.Error != nil {
		t.Errorf("Expected no error.  %s", response.Error)
	}
}
