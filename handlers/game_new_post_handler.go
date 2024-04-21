package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"solitaire/game"
)

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
			http.Error(w, "Error decoding game data", http.StatusBadRequest)
			return
	}

	newGame := game.NewGame()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newGame)
}

func postHandler(w http.ResponseWriter, r *http.Request) {


    // Read the request body
    requestBody, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    // Process the request body (replace this with your actual logic)
    fmt.Println("Received POST data:", string(requestBody))

    // Send a response
    fmt.Fprint(w, "POST request received!")
}
