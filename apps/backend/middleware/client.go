package middleware

import (
	"context"
	"net/http"

	"backend/utils"
)

type ClientAuthKey string

func ClientAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		token := authHeader[7:] // Strip "Bearer " prefix

		clientID, err := utils.ValidateClientJWT(token)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ClientIDKey, clientID)
		next(w, r.WithContext(ctx))
	}
}
