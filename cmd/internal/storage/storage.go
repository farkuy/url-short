package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	checkUrlTable  = "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'url');"
	createUrlTable = "CREATE TABLE IF NOT EXISTS url (id SERIAL PRIMARY KEY, alias TEXT UNIQUE NOT NULL, originalUrl TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"
)

func Start(path string) (*sql.DB, error) {
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, fmt.Errorf("Error conncect db: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Error ping db: %v", err)
	}

	var exists bool
	if err = db.QueryRow(checkUrlTable).Scan(&exists); err != nil {
		return nil, fmt.Errorf("Error check row: %v", err)
	}

	if !exists {
		if _, err = db.Exec(createUrlTable); err != nil {
			return nil, fmt.Errorf("Ошибка создания таблицы: %v", err)
		}
	}

	return db, nil
}
