package handlers

import (
	"encoding/json"
	"net/http"
)

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bc.Blocks)
}