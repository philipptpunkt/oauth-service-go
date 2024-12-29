package main

import (
	"log"
	"net/http"

	v1 "backend/handlers/v1"
	v1_clients "backend/handlers/v1/clients"
	"backend/middleware"
	"backend/models"
	"backend/utils"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	// Initialize Postgres
	utils.InitDatabase()
	db := utils.GetDB()

	err := db.AutoMigrate(&models.ClientCredential{}, &models.ClientRefreshToken{}, &models.OrganisationMember{}, &models.Organisation{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed")

	// Initialize Redis
	utils.InitRedis()

	// Routes General
	http.HandleFunc("/api/v1/health", v1.HealthHandler)

	// Routes Client
	http.HandleFunc("/api/v1/client/register", v1_clients.RegisterClientHandler)
	http.HandleFunc("/api/v1/client/login", v1_clients.LoginClientHandler)

	// Protected Client Routes
	http.Handle("/api/v1/client/verify-email-code", middleware.TemporaryAuthMiddleware(http.HandlerFunc(v1_clients.VerifyCodeHandler)))

	http.Handle("/api/v1/client/logout", middleware.ClientAuthMiddleware(v1_clients.LogoutClientHandler))
	http.Handle("/api/v1/client/delete", middleware.ClientAuthMiddleware(v1_clients.DeleteClientHandler))

	// Routes Auth Service

	log.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
