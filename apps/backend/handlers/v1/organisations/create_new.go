package v1_organisations

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"backend/middleware"
	"backend/utils"
)

type CreateOrganisationRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	LogoURL     string `json:"logo_url,omitempty"`
}

type CreateOrganisationResponse struct {
	Message        string `json:"message"`
	OrganisationID int    `json:"organisation_id"`
}

func CreateOrganisationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientID := r.Context().Value(middleware.ClientIDKey).(int) // From middleware
	var req CreateOrganisationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Organisation name is required", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()

	// Create organisation
	var organisationID int
	err := db.QueryRow(`
		INSERT INTO organisations (name, description, logo_url, owner_id)
		VALUES ($1, $2, $3, $4) RETURNING id
	`, req.Name, req.Description, req.LogoURL, clientID).Scan(&organisationID)
	if err != nil {
		log.Println("Error creating organisation:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Add owner as member
	_, err = db.Exec(`
		INSERT INTO organisation_members (organisation_id, client_id, role, created_at)
		VALUES ($1, $2, $3, $4)
	`, organisationID, clientID, "owner", time.Now())
	if err != nil {
		log.Println("Error adding organisation member:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateOrganisationResponse{
		Message:        "Organisation created successfully.",
		OrganisationID: organisationID,
	})
}
