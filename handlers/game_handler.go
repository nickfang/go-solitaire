package handlers

import (
	"net/http"
	"html/template"
	"encoding/json"
	"github.com/rs/zerolog/log"

	"solitaire/game"
)

var currentGame *game.Game

func GameHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().
        Str("method", r.Method).
        Str("url", r.URL.Path).
        Msg("Request received")
	if r.URL.Path != "/solitaire" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		json.NewEncoder(w).Encode("Welcome to the solitaire api.")
		tmpl, err := template.ParseFiles("templates/game.html")
		if err != nil {
				// Handle error
		}
		err = tmpl.Execute(w, currentGame) // Pass in your game data
		if err != nil {
				// Handle error
		}
	} else if r.Method == "POST" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}