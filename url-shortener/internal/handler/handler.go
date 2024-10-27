package handler

import (
	"encoding/json"
	"net/http"

	"example.com/url-shortener/internal/model"
	"example.com/url-shortener/internal/shortener"
	"example.com/url-shortener/internal/storage"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/shorten", ShortenURL)
	mux.HandleFunc("/u/", RedirectURL)
}

var store = storage.NewURLStore()

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input model.URL
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.OriginalURL == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	shortURL := shortener.GenerateShortURL()
	store.Save(shortURL, input.OriginalURL)

	response := map[string]string{"short_url": "https://short.ly/" + shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/u/"):]
	originalURL, exists := store.Find(shortURL)

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
