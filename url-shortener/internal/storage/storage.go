package storage

import "sync"

type URLStore struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string),
	}
}

func (s *URLStore) Save(shortURL, originalURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortURL] = originalURL
}

func (s *URLStore) Find(shortURL string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	originalURL, exists := s.urls[shortURL]
	return originalURL, exists
}
