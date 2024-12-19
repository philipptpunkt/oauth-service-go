package main

import (
	v1 "backend/handlers/v1"
	v1_clients "backend/handlers/v1/clients"
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

	// Routes General
	http.HandleFunc("/api/v1/health", v1.HealthHandler)

	// Routes Client
	http.HandleFunc("/api/v1/client/register", v1_clients.RegisterClientHandler)

	// Routes Auth Service

	log.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}