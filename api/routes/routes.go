package routes

import (
	"net/http"
	"solitaire/api/handlers"
)

func Routes() {
	http.HandleFunc("/solitaire", handlers.GameHandler)
	http.HandleFunc("/solitaire/new", handlers.CreateHandler)
	http.HandleFunc("/solitaire/info", handlers.InfoHandler)
	http.HandleFunc("/solitaire/move", handlers.MoveHandler)
	http.HandleFunc("/solitaire/hint", handlers.HintHandler)
	http.HandleFunc("/solitaire/reset", handlers.ResetHandler)
}
