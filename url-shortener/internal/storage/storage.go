package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type URLStore struct {
	db *sql.DB
}

func NewURLStore(dataSourceName string) *URLStore {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (
		short_url VARCHAR(10) PRIMARY KEY,
		original_url TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatalf("could not create table: %v", err)
	}

	return &URLStore{db: db}

}

func (s *URLStore) Save(shortURL, originalURL string) {
	_, err := s.db.Exec("INSERT INTO urls (short_url, original_url) VALUES (?, ?)", shortURL, originalURL)
	if err != nil {
		log.Fatalf("Could not insert data: %v", err)
	}
	// id, err := res.LastInsertId()
	// return id
}

func (s *URLStore) Find(shortURL string) (string, bool) {
	var originalURL string

	err := s.db.QueryRow("SELECT original_url FROM urls WHERE short_url = ?", shortURL).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false
		}
		log.Printf("Could not find URL: %v", err)
		return "", false
	}

	return originalURL, true
}
