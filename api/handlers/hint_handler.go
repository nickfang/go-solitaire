package handlers

import (
	"encoding/json"
	"net/http"
)

func HintHandler(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	if r.Method == "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Under Construction.")
}
