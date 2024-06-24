package main

import (
	"fmt"
	"strings"

	"solitaire/gamemanager"
)

func main() {
	gm := gamemanager.NewGameManager()
	go gm.ProcessRequests()

	gameId, error := gm.CreateSession()
	if error != nil {
		fmt.Println(error)
		return
	}
	session, err := gm.GetSession(gameId)
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
		gr := gamemanager.GameRequest{SessionId: gameId, Action: input, Response: responseChan}
		gm.Requests <- gr

		response := <-responseChan
		if response.Error != nil {
			fmt.Println(response.Error)
		}
		if input != "ss" && input != "h" && input != "?" {
			DisplayGame(*response.Game)
		}

	}
}
