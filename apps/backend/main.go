package main

import (
	v1 "backend/handlers/v1"
	"backend/utils"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	
	// Initialize Postgres
	utils.InitDatabase()
	// Initialize Redis
	utils.InitRedis()

	// Routes and Request Handlers
	http.HandleFunc("/api/v1/health", v1.HealthHandler)

	log.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}