package v1_organisations

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/middleware"
	"backend/models"
	"backend/utils"

	"gorm.io/gorm"
)

type CreateOrganisationRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	LogoURL     string `json:"logo_url,omitempty"`
}

type CreateOrganisationResponse struct {
	Message        string `json:"message"`
	OrganisationID uint   `json:"organisation_id"`
}

func CreateOrganisationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientID := r.Context().Value(middleware.ClientIDKey).(uint)
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

	err := db.Transaction(func(tx *gorm.DB) error {
		organisation := models.Organisation{
			Name:        req.Name,
			Description: req.Description,
			LogoURL:     req.LogoURL,
			OwnerID:     clientID,
		}
		if err := tx.Create(&organisation).Error; err != nil {
			return err
		}

		member := models.OrganisationMember{
			OrganisationID: organisation.ID,
			ClientID:       clientID,
			Role:           models.RoleOwner,
			JoinedAt:       organisation.CreatedAt,
		}
		if err := tx.Create(&member).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Println("Error creating organisation or member:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateOrganisationResponse{
		Message:        "Organisation created successfully.",
		OrganisationID: clientID,
	})
}
