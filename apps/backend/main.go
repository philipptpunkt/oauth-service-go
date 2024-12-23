package main

import (
	"log"
	"net/http"

	v1 "backend/handlers/v1"
	v1_clients "backend/handlers/v1/clients"
	"backend/middleware"
	"backend/utils"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	// Initialize Postgres
	utils.InitDatabase()
	// Initialize Redis
	utils.InitRedis()

	// Routes General
	http.HandleFunc("/api/v1/health", v1.HealthHandler)
	http.Handle("/api/v1/cookie-test", middleware.CORSMiddleware(http.HandlerFunc(v1.TestCookieHandler)))

	// Routes Client
	http.HandleFunc("/api/v1/client/register", v1_clients.RegisterClientHandler)

	// Protected routes
	http.Handle("/api/v1/client/verify-email-code", middleware.TemporaryAuthMiddleware(http.HandlerFunc(v1_clients.VerifyCodeHandler)))

	// Routes Auth Service

	log.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
