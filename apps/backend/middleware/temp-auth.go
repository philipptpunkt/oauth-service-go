package middleware

import (
	"context"
	"net/http"

	"backend/utils"
)

type TemporaryAuthKey string

const (
	ClientIDKey TemporaryAuthKey = "clientID"
	PurposeKey  TemporaryAuthKey = "purpose"
)

func TemporaryAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		token := authHeader[7:] // Strip "Bearer " prefix

		clientID, purpose, err := utils.ValidateTemporaryJWT(token)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ClientIDKey, clientID)
		ctx = context.WithValue(ctx, PurposeKey, purpose)
		next(w, r.WithContext(ctx))
	}
}
