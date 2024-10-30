package main

import (
	"log"
	"net/http"

	"example.com/url-shortener/config"
	"example.com/url-shortener/internal/handler"
	"example.com/url-shortener/internal/storage"
)

func main() {
	cfg := config.LoadConfig()
	dataSourceName := cfg.DBUSER + ":" + cfg.DBPASSWORD + "@tcp(" + cfg.DBHOST + ":" + cfg.DBPORT + ")/" + cfg.DBNAME + "?allowNativePasswords=true"

	store := storage.NewURLStore(dataSourceName)
	handler.SetStore(store)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start the server: %s\n", err.Error())
	}
}
