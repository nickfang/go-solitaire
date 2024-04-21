package main

import (
	"fmt"
	"html/template"
	"net/http"
	"solitaire/game"
	// "solitaire/gamemanagerapi"
	"solitaire/handlers"
)

var currentGame *game.Game

func renderGame(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/game.html")
	if err != nil {
			// Handle error
	}
	err = tmpl.Execute(w, currentGame) // Pass in your game data
	if err != nil {
			// Handle error
	}
}

func handleGameAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleGameAction")
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		renderGame(w)
	} else if r.Method == "POST" {
		handleGameAction(w, r)
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", gameHandler)
	http.HandleFunc("/solitaire/new", handlers.CreateGameHandler)
	http.ListenAndServe(":8888", nil)
}