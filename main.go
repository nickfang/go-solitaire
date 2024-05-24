package main

import (
	"net/http"

	"github.com/rs/zerolog"

	"solitaire/api"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	api.StartApi()

	http.ListenAndServe(":8888", nil)
}
