package database

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/joeCavZero/blogland/logger"
	_ "github.com/lib/pq"
)

var DEFAULT_SESSION_TOKENS_COLLECTOR_SECONDS int64 = 60

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

	// this is made to free tokens that are older than x time
	go func() {
		sessionTokensCollectorSecondsString := os.Getenv("SESSION_TOKENS_COLLECTOR_SECONDS")
		sessionTokensCollectorSeconds := DEFAULT_SESSION_TOKENS_COLLECTOR_SECONDS
		if sessionTokensCollectorSecondsString == "" {
			logger.APIInfof("SESSION_TOKENS_COLLECTOR_SECONDS not set, using default: %d", DEFAULT_SESSION_TOKENS_COLLECTOR_SECONDS)
		} else {
			res, err := strconv.Atoi(sessionTokensCollectorSecondsString)
			if err != nil {
				logger.APIExitErrorf("Invalid SESSION_TOKENS_COLLECTOR_SECONDS: %v", err)
			}

			sessionTokensCollectorSeconds = int64(res)

		}

		dt := New(Database)
		ctx := context.Background()
		for {
			err := dt.DeleteOldSessionTokens(ctx, sessionTokensCollectorSeconds)
			if err != nil {
				logger.APIErrorf("Failed to delete old tokens: %v", err)
			}

			time.Sleep(15 * time.Second)
		}
	}()
}

func startDatabaseTables() {
	dt := New(Database)
	ctx := context.Background()

	dt.CreateUsersTable(ctx)
	logger.APIInfof("Users table created or already exists")

	dt.CreateSessionTokensTable(ctx)
	logger.APIInfof("Tokens table created or already exists")
}
