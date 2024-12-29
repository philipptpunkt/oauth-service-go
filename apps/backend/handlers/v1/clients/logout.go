package v1_clients

import (
	"log"
	"net/http"

	"backend/middleware"
	"backend/models"
	"backend/utils"
)

func LogoutClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientID := r.Context().Value(middleware.ClientIDKey).(int)

	db := utils.GetDB()

	if err := db.Where("client_id = ?", clientID).Delete(&models.ClientRefreshToken{}).Error; err != nil {
		log.Println("Error removing refresh tokens:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
