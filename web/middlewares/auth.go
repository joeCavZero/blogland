package middlewares

import (
	"net/http"
	"strconv"

	"github.com/joeCavZero/blogland/logger"
	"github.com/joeCavZero/blogland/web/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sessionTokenCookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Unauthorized: Missing session token", http.StatusUnauthorized)
			return
		}

		sessionUserIDCookie, err := r.Cookie("session_user_id")
		if err != nil {
			logger.WebErrorf("Error retrieving session user ID cookie: %v", err)
			http.Error(w, "Unauthorized: Missing session user ID", http.StatusUnauthorized)
			return
		}

		res, err := strconv.Atoi(sessionUserIDCookie.Value)
		if err != nil {
			http.Error(w, "Invalid session user ID", http.StatusUnauthorized)
			return
		}

		sessionToken := sessionTokenCookie.Value
		sessionUserID := int64(res)

		if !auth.ValidateSessionToken(sessionUserID, sessionToken) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
