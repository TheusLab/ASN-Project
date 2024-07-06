package main

import (
	"log"
	"net/http"

	"github.com/TheusLab/ASN-Project/backend/handlers"
)

func main() {
	http.HandleFunc("/api/search", handlers.SearchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
