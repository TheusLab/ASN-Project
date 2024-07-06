package main

import (
	"asn/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/search", handlers.SearchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
