package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"

	"solitaire/game"
	"solitaire/game/deck"
	"solitaire/game/stacks"
	// "solitaire/game/gamestates"
)

var currentGame *game.Game

type ResponseData struct {
	Message string `json:"message"`
	Data 		any		 `json:"data"`
}

type DeckUser struct {
	Card deck.Card `json:"card"`
	NumBeforeCard int `json:"numBeforeCard"`
	NumAfterCard int `json:"numAfterCard"`
}
type StacksUser struct {
	Stacks stacks.Stacks `json:"topStackCards"`
}
type BoardUser struct {
	Column1Cards []string `json:"column1Cards"`
	Column1NumHidden int `json:"column1NumHidden"`
	Column2Cards []string `json:"column2Cards"`
	Column2NumHidden int `json:"column2NumHidden"`
	Column3Cards []string `json:"column3Cards"`
	Column3NumHidden int `json:"column3NumHidden"`
	Column4Cards []string `json:"column4Cards"`
	Column4NumHidden int `json:"column4NumHidden"`
	Column5Cards []string `json:"column5Cards"`
	Column5NumHidden int `json:"column5NumHidden"`
	Column6Cards []string `json:"column6Cards"`
	Column6NumHidden int `json:"column6NumHidden"`
	Column7Cards []string `json:"column7Cards"`
	Column7NumHidden int `json:"column7NumHidden"`
}
type GameUser struct {
	Deck DeckUser
	Stacks StacksUser
	Board BoardUser
}

func LogRequest(r *http.Request) {
	log.Info().
		Str("method", r.Method).
		Str("url", r.URL.Path).
		Msg("Request received")
}

func LogResponse(w http.ResponseWriter) {
	// TODO: Handle different status and errors
	log.Info().
		Int("status", 200).
		Msg("Move completed successfully")
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
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

