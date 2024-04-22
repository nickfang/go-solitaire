package main

import (
	"net/http"
	"github.com/rs/zerolog"

	"solitaire/handlers"
)



func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	http.HandleFunc("/solitaire", handlers.GameHandler)
	http.HandleFunc("/solitaire/new", handlers.CreateHandler)
	http.HandleFunc("/solitaire/info", handlers.InfoHandler)
	http.HandleFunc("/solitaire/move", handlers.MoveHandler)
	http.HandleFunc("/solitaire/hint", handlers.HintHandler)
	http.HandleFunc("/solitaire/reset", handlers.ResetHandler)
	http.ListenAndServe(":8888", nil)
}