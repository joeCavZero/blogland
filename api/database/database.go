package database

import (
	"database/sql"
	"os"

	"github.com/joeCavZero/blogland/logger"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func StartDatabase() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logger.APIExitErrorf("Database URL not set in environment variables")
	}

	Database, err = sql.Open("postgres", dbURL)
	if err != nil {
		logger.APIExitErrorf("Failed to connect to database: %v", err)
	}

	startDatabaseTables()
}

func startDatabaseTables() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(64) UNIQUE NOT NULL,
			password VARCHAR(64) NOT NULL,
			role VARCHAR(16)
		);
	`
	_, err := Database.Exec(query)
	if err != nil {
		logger.APIExitErrorf("Failed to create users table: %v", err)
	}
	logger.APIInfof("Users table created or already exists")

	query = `
		CREATE TABLE IF NOT EXISTS tokens (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			token VARCHAR(256) UNIQUE NOT NULL,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err = Database.Exec(query)
	if err != nil {
		logger.APIExitErrorf("Failed to create tokens table: %v", err)
	}
	logger.APIInfof("Tokens table created or already exists")
}
