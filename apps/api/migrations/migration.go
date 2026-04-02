package database

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {

	query := `
	CREATE TABLE IF NOT EXISTS users (
	    id SERIAL PRIMARY KEY,
	    username VARCHAR(50) UNIQUE NOT NULL,
	    password TEXT NOT NULL,
	    role VARCHAR(20),
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS vehicles (
		id SERIAL PRIMARY KEY,
		license_plate VARCHAR(20) UNIQUE NOT NULL,
		tariff INTEGER NOT NULL,
		duration INTERVAL,
		owner_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration success")
}
