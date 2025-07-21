package auth

import (
	"context"

	"github.com/joeCavZero/blogland/api/database"
)

func ValidateSessionToken(sessionUserId int64, sessionToken string) bool {
	dt := database.New(database.Database)
	ctx := context.Background()
	_, err := dt.GetSessionTokenByIDAndToken(
		ctx,
		database.GetSessionTokenByIDAndTokenParams{
			UserID: sessionUserId,
			Token:  sessionToken,
		},
	)
	return err == nil
}
