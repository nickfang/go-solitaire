package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solitaire/solitairestore"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	if r.Method == "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	queryValues := r.URL.Query()
	gameId := queryValues.Get("game-id")
	if gameId == "" {
		response := ResponseData{
			Message: "Error: missing game-id.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	game, gamestates, err := solitairestore.New().LoadGame(gameId)
	if err != nil {
		fmt.Println(err.Error())
	}
	gameResponse := ResponseData{"Game found.", GameData{game, gamestates}}
	fmt.Println(gameResponse)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gameResponse)
}
