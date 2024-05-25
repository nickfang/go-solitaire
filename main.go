package main

import (
	"net/http"

	"github.com/rs/zerolog"

	"solitaire/api"
	"solitaire/game"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	api.StartApi()
	game.NewGame()

	http.ListenAndServe(":8888", nil)
}
