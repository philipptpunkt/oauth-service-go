package v1_clients

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/middleware"
	"backend/utils"
)

type VerifyCodeRequest struct {
	Code string `json:"code"`
}

type VerifyCodeResponse struct {
	Message string `json:"message"`
}

func VerifyCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.Context().Value(middleware.EmailKey).(string)
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

	err := utils.ValidateVerificationCode(ctx, rdb, email, req.Code)
	if err != nil {
		log.Printf("Verification failed for email %s: %v", email, err)
		http.Error(w, "Invalid or expired code", http.StatusUnauthorized)
		return
	}

	db := utils.GetDB()
	_, err = db.Exec("UPDATE client_credentials SET email_verified = TRUE WHERE email = $1", email)
	if err != nil {
		log.Println("Error updating email verification status:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VerifyCodeResponse{
		Message: "Email successfully verified.",
	})
}
