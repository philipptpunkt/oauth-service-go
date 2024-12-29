package v1_clients

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/middleware"
	"backend/models"
	"backend/utils"
)

type UpdateClientRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Organisation   string `json:"organisation"`
	JobTitle       string `json:"job_title"`
	ProfilePicture string `json:"profile_picture"`
	TimeZone       string `json:"time_zone"`
}

func ClientProfileHandler(w http.ResponseWriter, r *http.Request) {
	clientID := r.Context().Value(middleware.ClientIDKey).(int)

	db := utils.GetDB()

	switch r.Method {
	case http.MethodGet:
		var client models.ClientProfile
		err := db.Where("client_credentials_id = ?", clientID).First(&client).Error
		if err != nil {
			log.Println("Error fetching client data:", err)
			http.Error(w, "Client not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(client)

	case http.MethodPut:
		var req UpdateClientRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var client models.ClientProfile
		err := db.Where("client_credentials_id = ?", clientID).First(&client).Error
		if err != nil {
			log.Println("Error fetching client data:", err)
			http.Error(w, "Client not found", http.StatusNotFound)
			return
		}

		client.FirstName = req.FirstName
		client.LastName = req.LastName
		client.Organisation = req.Organisation
		client.JobTitle = req.JobTitle
		client.ProfilePicture = req.ProfilePicture
		client.TimeZone = req.TimeZone

		err = db.Save(&client).Error
		if err != nil {
			log.Println("Error updating client data:", err)
			http.Error(w, "Failed to update client data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Client data updated successfully"))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
