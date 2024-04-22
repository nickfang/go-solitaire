package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().
		Str("method", r.Method).
		Str("url", r.URL.Path).
		Msg("Request received")
	if r.Method == "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Under Construction.")
}