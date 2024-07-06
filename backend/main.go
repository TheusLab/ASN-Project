package main

import (
	"encoding/json"
	"net/http"
	"github.com/TheusLab/ASN-Project/backend/utils"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := utils.Search(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
