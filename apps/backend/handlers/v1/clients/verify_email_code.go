package v1_clients

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"backend/middleware"
	"backend/utils"
)

type VerifyCodeRequest struct {
	Code string `json:"code"`
}

type VerifyCodeResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func VerifyCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientID := r.Context().Value(middleware.ClientIDKey).(int)
	purpose := r.Context().Value(middleware.PurposeKey).(string)

	if purpose != "email_verification" {
		http.Error(w, "Unauthorized: invalid token purpose", http.StatusUnauthorized)
		return
	}

	var req VerifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Code == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	rdb := utils.GetRedisClient()
	ctx := utils.GetRedisContext()

	err := utils.ValidateVerificationCode(ctx, rdb, clientID, req.Code)
	if err != nil {
		log.Printf("Verification failed for client ID %d: %v", clientID, err)
		http.Error(w, "Invalid or expired code", http.StatusUnauthorized)
		return
	}

	db := utils.GetDB()

	_, err = db.Exec("UPDATE client_credentials SET email_verified = TRUE WHERE id = $1", clientID)
	if err != nil {
		log.Println("Error updating email verification status:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	accessToken, err := utils.GenerateClientJWT(clientID, utils.RoleOwner, 15*time.Minute)
	if err != nil {
		log.Println("Error generating access token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		log.Println("Error generating refresh token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	encryptedToken, err := utils.EncryptToken(refreshToken, utils.GetEncryptionKey())
	if err != nil {
		log.Println("Error encrypting refresh token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour) // 7 days
	_, err = db.Exec(
		"INSERT INTO client_refresh_tokens (client_id, token, expires_at) VALUES ($1, $2, $3)",
		clientID, encryptedToken, expiresAt,
	)
	if err != nil {
		log.Println("Error storing refresh token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VerifyCodeResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}
