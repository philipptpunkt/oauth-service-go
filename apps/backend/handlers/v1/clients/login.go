package v1_clients

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"backend/models"
	"backend/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginClientRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginClientResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
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

	var client models.ClientCredentials
	err := db.Where("email = ?", req.Email).First(&client).Error
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Println("Error querying client credentials:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateClientJWT(int(client.ID), 15*time.Minute)
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
	refreshTokenEntry := models.ClientRefreshToken{
		ClientID:  client.ID,
		Token:     encryptedToken,
		ExpiresAt: expiresAt,
	}
	if err := db.Create(&refreshTokenEntry).Error; err != nil {
		log.Println("Error storing refresh token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginClientResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}
