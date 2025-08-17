package storage

import (
	"database/sql"
	"fmt"
)

const (
	ckeckUrlTable  = "SELECT * FROM information_schema.tables WHERE table_name = 'url';"
	createUrlTable = "CREATE TABLE IF NOT EXISTS url (id SERIAL PRIMARY KEY, alias TEXT UNIQUE NOT NULL, originalUrl TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"
)

func Start(path string) (*sql.DB, error) {
	db, error := sql.Open("postgres", path)
	if error != nil {
		return nil, fmt.Errorf("Error conncect db: %v", error)
	}
	if error = db.Ping(); error != nil {
		return nil, fmt.Errorf("Error ping db: %v", error)
	}

	var exists bool
	error = db.QueryRow(ckeckUrlTable).Scan(exists)
	if error != nil {
		return nil, fmt.Errorf("Error check row: %v", error)
	}

	if !exists {
		_, error := db.Exec(createUrlTable)

		if error != nil {
			return nil, fmt.Errorf("Error create row: %v", error)
		}
	}

	return db, nil
}
