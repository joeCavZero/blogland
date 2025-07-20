package models

import (
	"net/http"
	"time"
)

func CreateTokenCookie(token string) *http.Cookie {
	expires := time.Now().Add(24 * time.Hour)

	return &http.Cookie{
		Name:    "auth_token",
		Value:   token,
		Expires: expires,
	}
}
