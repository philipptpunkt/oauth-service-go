package v1_clients

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginClientRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginClientResponse struct {
	Token            string `json:"token"`
	RefreshToken     string `json:"refresh_token"`
	ProfileCompleted bool   `json:"profile_completed"`
}

func LoginClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Email == "" || req.Password == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()

	var hashedPassword string
	var clientID int
	var profile_completed bool
	err := db.QueryRow("SELECT id, password, profile_completed FROM client_credentials WHERE email = $1", req.Email).Scan(&clientID, &hashedPassword, &profile_completed)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Println("Error querying client credentials:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateClientJWT(clientID, 15*time.Minute)
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

	encryptionKey, err := utils.GetEncryptionKey()
	if err != nil {
		log.Println("Error fetching encryption key:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	encryptedToken, err := utils.EncryptToken(refreshToken, encryptionKey)
	if err != nil {
		log.Println("Error encrypting refresh token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	_, err = db.Exec("INSERT INTO client_refresh_tokens (client_id, token, expires_at) VALUES ($1, $2, $3)", clientID, encryptedToken, expiresAt)
	if err != nil {
		log.Println("Error storing refresh token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginClientResponse{
		Token:            accessToken,
		RefreshToken:     refreshToken,
		ProfileCompleted: profile_completed,
	})
}
