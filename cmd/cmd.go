package main

import (
	"fmt"
	"strings"

	"solitaire/gamemanager"
)

func main() {
	gm := gamemanager.NewGameManager()
	go gm.GameEngine()
	go gm.SessionEngine()

	gm.SessionReq <- gamemanager.SessionRequest{Action: "create"}
	sessionRes := <-gm.SessionRes
	sessionId := sessionRes.Id
	gm.GameReq <- gamemanager.GameRequest{SessionId: sessionId, Action: "display"}
	gameRes := <-gm.GameRes
	game := gameRes.Game

	DisplayGame(*game)
	var i string

	for {
		fmt.Scanln(&i)
		input := strings.ToLower(i)
		gm.GameReq <- gamemanager.GameRequest{SessionId: sessionId, Action: "display"}
		gameRes := <-gm.GameRes
		game = gameRes.Game
		if input == "r" {
			gm.SessionReq <- gamemanager.SessionRequest{Action: "create"}
			sessionRes = <-gm.SessionRes
			if sessionRes.Error != nil {
				fmt.Println(sessionRes.Error)
				continue
			}
			sessionId = sessionRes.Id
			gm.GameReq <- gamemanager.GameRequest{SessionId: sessionId, Action: "display"}
			gameRes = <-gm.GameRes
			game = gameRes.Game
			DisplayGame(*game)
			continue
		}
		if input == "rt" {
			gm.SessionReq <- gamemanager.SessionRequest{Id: sessionId, Action: "delete"}
			session := <-gm.SessionRes
			if session.Error != nil {
				fmt.Println(session.Error)
				continue
			}

			gm.SessionReq <- gamemanager.SessionRequest{Action: "create:test"}
			session = <-gm.SessionRes
			if session.Error != nil {
				fmt.Println(session.Error)
				continue
			}
			sessionId = session.Id
			gm.GameReq <- gamemanager.GameRequest{SessionId: sessionId, Action: "display"}
			gameRes = <-gm.GameRes
			game = gameRes.Game
			DisplayGame(*game)
			continue
		}
		gr := gamemanager.GameRequest{SessionId: sessionId, Action: input}
		gm.GameReq <- gr

		gameRes = <-gm.GameRes
		if gameRes.Error != nil {
			if gameRes.Error.Error() == "quit" {
				fmt.Println("Quitting...")
				break
			}
			if gameRes.Error.Error() == "finished" {
				fmt.Println("Congrats! You won!")
				gm.SessionReq <- gamemanager.SessionRequest{Id: sessionId, Action: "delete"}
				session := <-gm.SessionRes
				gm.SessionReq <- gamemanager.SessionRequest{Action: "create"}
				session = <-gm.SessionRes
				sessionId = session.Id
				gm.GameReq <- gamemanager.GameRequest{SessionId: sessionId, Action: "display"}
				gameRes = <-gm.GameRes
				game = gameRes.Game
				DisplayGame(*game)
				continue
			}
			fmt.Println(gameRes.Error)
		}
		if input == "h" {
			DisplayHints(*game)
		} else if input == "?" {
			DisplayHelp()
		} else {
			DisplayGame(*game)
		}
	}
	gamemanager.CloseManager(gm)
}
