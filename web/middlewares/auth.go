package middlewares

import (
	"net/http"
	"strconv"

	"github.com/joeCavZero/blogland/logger"
	"github.com/joeCavZero/blogland/web/auth"
	"github.com/joeCavZero/blogland/web/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionTokenCookie, err := r.Cookie("session_token")
		if err != nil {
			utils.RenderErrorTemplate(w, http.StatusUnauthorized)
			return
		}

		sessionUserIDCookie, err := r.Cookie("session_user_id")
		if err != nil {
			utils.RenderErrorTemplate(w, http.StatusUnauthorized)
			logger.WebErrorf("Error retrieving session user ID cookie: %v", err)
			return
		}

		res, err := strconv.Atoi(sessionUserIDCookie.Value)
		if err != nil {
			utils.RenderErrorTemplate(w, http.StatusUnauthorized)
			return
		}

		sessionToken := sessionTokenCookie.Value
		sessionUserID := int64(res)

		if !auth.ValidateSessionToken(sessionUserID, sessionToken) {
			utils.RenderErrorTemplate(w, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
