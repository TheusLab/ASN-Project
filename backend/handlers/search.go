package handlers

import (
	"asn/utils"
	"encoding/json"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := utils.Search(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
