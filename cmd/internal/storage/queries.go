package storage

const (
	CheckUrlTable = `
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_schema = 'public'
              AND table_name = 'url'
        );
    `
	CreateUrlTable = `
        CREATE TABLE IF NOT EXISTS url (
            id SERIAL PRIMARY KEY,
            alias TEXT UNIQUE NOT NULL,
            originalUrl TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `
	PostUrlRow   = "INSERT INTO url (alias, originalUrl) VALUES ($1, $2)"
	GetUrlRow    = "SELECT originalUrl FROM url WHERE alias=$1"
	UpdateUrlRow = "UPDATE url SET originalUrl=$1 WHERE alias=$2"
	DeleteUrlRow = "DELETE FROM url WHERE alias=$1"
)
