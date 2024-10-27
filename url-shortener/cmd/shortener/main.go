package main

import (
	"log"
	"net/http"

	"example.com/url-shortener/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start the server: %s\n", err.Error())
	}
}
