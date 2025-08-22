package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func Start(path string) (*Storage, error) {
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка проверки соединения базы данных: %v", err)
	}

	var exists bool
	if err = db.QueryRow(checkUrlTable).Scan(&exists); err != nil {
		return nil, fmt.Errorf("ошибка проверки существования таблицы: %v", err)
	}

	if !exists {
		if _, err = db.Exec(createUrlTable); err != nil {
			return nil, fmt.Errorf("ошибка создания таблицы: %v", err)
		}
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUrl(alias, longUrl string) error {
	_, err := s.db.Exec(postUrlRow, alias, longUrl)
	if err != nil {
		var pqErr *pq.Error
		errors.As(err, &pqErr)

		switch pqErr.Code.Name() {
		case "unique_violation":
			return fmt.Errorf("alias уже занят: %v", alias)
		default:
			return fmt.Errorf("ошибка добавления url (%v) и alias (%v): %v", longUrl, alias, err)
		}
	}

	return nil
}

func (s *Storage) GetUrl(alias string) (string, error) {
	var originalUrl string
	err := s.db.QueryRow(getUrlRow, alias).Scan(&originalUrl)
	if err != nil {
		return "", fmt.Errorf("ошибка получения alias (%v): %v", alias, err)
	}

	return originalUrl, nil
}

func (s *Storage) UpdateUrl(alias, newUrl string) error {
	res, err := s.db.Exec(updateUrlRow, newUrl, alias)
	if err != nil {
		return fmt.Errorf("ошибка обновления alias (%v): %v", alias, err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("alias не найден: %v", alias)
	}

	return nil
}

// TODO: написать проверку существования alias
func (s *Storage) DeleteUrl(alias string) (string, error) {
	var deletedUrl string
	err := s.db.QueryRow(deleteUrlRow, alias).Scan(&deletedUrl)
	if err != nil {
		return "", fmt.Errorf("ошибка удаления алиаса (%v): %v", alias, err)
	}

	return deletedUrl, nil
}
