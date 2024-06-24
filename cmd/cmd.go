package main

import (
	"fmt"
	"strings"

	"solitaire/gamemanager"
)

func main() {
	gm := gamemanager.NewGameManager()
	go gm.ProcessRequests()

	sessionId, error := gm.CreateSession()
	if error != nil {
		fmt.Println(error)
		return
	}

	error = gm.InitializeGame(sessionId)
	if error != nil {
		fmt.Println(error)
		return
	}

	session, err := gm.GetSession(sessionId)
	if err != nil {
		fmt.Println(err)
		return
	}

	game := session.Game
	DisplayGame(*game)
	var i string
	responseChan := make(chan gamemanager.GameResponse)

	for {
		fmt.Scanln(&i)
		input := strings.ToLower(i)
		if input == "r" {
			gm.DeleteSession(sessionId)
			sessionId, _ = gm.CreateSession()
			gm.InitializeGame(sessionId)
			session, _ = gm.GetSession(sessionId)
			DisplayGame(*session.Game)
			continue
		}
		if input == "rt" {
			gm.DeleteSession(sessionId)
			sessionId, _ = gm.CreateSession()
			gm.InitializeTestGame(sessionId)
			session, _ = gm.GetSession(sessionId)
			DisplayGame(*session.Game)
			continue
		}
		gr := gamemanager.GameRequest{SessionId: sessionId, Action: input, Response: responseChan}
		gm.Requests <- gr

		response := <-responseChan
		if response.Error != nil {
			if response.Error.Error() == "quit" {
				fmt.Println("Quitting...")
				return
			}
			if response.Error.Error() == "finished" {
				fmt.Println("Congrats! You won!")
				continue
			}
			fmt.Println(response.Error)
		}
		if input == "h" {
			DisplayHints(*session.Game)
		} else if input == "?" {
			DisplayHelp()
		} else {
			DisplayGame(*session.Game)
		}

	}
}
