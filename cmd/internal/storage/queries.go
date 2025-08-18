package storage

const (
	checkUrlTable = `
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_schema = 'public'
              AND table_name = 'url'
        );
    `
	createUrlTable = `
        CREATE TABLE IF NOT EXISTS url (
            id SERIAL PRIMARY KEY,
            alias TEXT UNIQUE NOT NULL,
            originalUrl TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `
	postUrlRow   = "INSERT INTO url (alias, originalUrl) VALUES ($1, $2)"
	getUrlRow    = "SELECT originalUrl FROM url WHERE alias=$1"
	updateUrlRow = "UPDATE url SET originalUrl=$1 WHERE alias=$2"
	deleteUrlRow = "DELETE FROM url WHERE alias=$1"
)
