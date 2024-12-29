package v1_clients

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"backend/models"
	"backend/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterClientRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterClientResponse struct {
	Token string `json:"token"`
}

func RegisterClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()

	var existingClient models.ClientCredential
	if err := db.Where("email = ?", req.Email).First(&existingClient).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println("Error checking email existence:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Email already registered", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	client := models.ClientCredential{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := db.Create(&client).Error; err != nil {
		log.Println("Error inserting client credentials:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	verificationCode := generateVerificationCode()

	redisClient := utils.GetRedisClient()
	redisCtx := utils.GetRedisContext()
	err = utils.StoreVerificationCode(redisCtx, redisClient, int(client.ID), verificationCode, time.Hour)
	if err != nil {
		log.Println("Error storing verification code in Redis:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	templatePath := "./templates/email_code_verification_client.html"

	data := map[string]interface{}{
		"Code":     verificationCode,
		"CopyLink": '/',
	}

	emailBody, err := utils.ParseHtmlTemplate(templatePath, data)
	if err != nil {
		log.Printf("Error parsing email template: %v\n", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	emailSender, _ := utils.CreateEmailSender()
	err = emailSender.SendEmail(req.Email, "Your Verification Code", emailBody, false)
	if err != nil {
		log.Println("Error sending email:", err)
		http.Error(w, "Failed to send verification email", http.StatusInternalServerError)
		return
	}

	tempToken, err := utils.GenerateTemporaryJWT(int(client.ID), "email_verification", time.Hour)
	if err != nil {
		log.Println("Error generating temporary token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RegisterClientResponse{
		Token: tempToken,
	})
}

func generateVerificationCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
