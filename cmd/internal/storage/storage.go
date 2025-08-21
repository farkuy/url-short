package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func Start(path string) (*Storage, error) {
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, fmt.Errorf("error conncect db. %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error ping db. %v", err)
	}

	var exists bool
	if err = db.QueryRow(checkUrlTable).Scan(&exists); err != nil {
		return nil, fmt.Errorf("error checking table existence. %v", err)
	}

	if !exists {
		if _, err = db.Exec(createUrlTable); err != nil {
			return nil, fmt.Errorf("error creating table. %v", err)
		}
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUrl(alias, longUrl string) error {
	_, err := s.db.Exec(postUrlRow, alias, longUrl)
	if err != nil {
		return fmt.Errorf("error adding URL (%v) and alias (%v). (%v)", longUrl, alias, err)
	}

	return nil
}

func (s *Storage) GetUrl(alias string) (string, error) {
	var originalUrl string
	err := s.db.QueryRow(getUrlRow, alias).Scan(&originalUrl)
	if err != nil {
		return "", fmt.Errorf("error get alias (%v). (%v)", alias, err)
	}

	return originalUrl, nil
}

func (s *Storage) UpdateUrl(alias, newUrl string) error {
	_, err := s.db.Exec(updateUrlRow, newUrl, alias)
	if err != nil {
		return fmt.Errorf("error update alias (%v). %v", alias, err)
	}

	return nil
}

func (s *Storage) DeleteUrl(alias string) (string, error) {
	var deletedUrl string
	err := s.db.QueryRow(deleteUrlRow, alias).Scan(&deletedUrl)
	if err != nil {
		return "", fmt.Errorf("error deleting alias (%v). %v", alias, err)
	}

	return deletedUrl, nil
}
