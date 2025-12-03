package middlewares

import (
	"context"
	"net/http"
	"strings"
	"twitter-clone/utils"
)

type contextKey struct{}

var UserIDKey = contextKey{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		var tokenString string

		if token, ok := strings.CutPrefix(authHeader, "Bearer "); ok {
			tokenString = token
		}

		if tokenString == "" {
			if cookie, err := r.Cookie("access_token"); err == nil {
				tokenString = cookie.Value
			}
		}

		if tokenString == "" {
			utils.Error(w, http.StatusUnauthorized, "authentication required")
			return
		}

		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.Error(w, http.StatusUnauthorized, "invalid or expired token", err)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
