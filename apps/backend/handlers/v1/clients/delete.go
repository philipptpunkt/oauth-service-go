package v1_clients

import (
	"log"
	"net/http"

	"backend/middleware"
	"backend/models"
	"backend/utils"
)

func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientID := r.Context().Value(middleware.ClientIDKey).(int)

	db := utils.GetDB()

	tx := db.Begin()
	if err := tx.Error; err != nil {
		log.Println("Error starting transaction:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if err := tx.Where("client_id = ?", clientID).Delete(&models.ClientRefreshToken{}).Error; err != nil {
		log.Println("Error deleting refresh tokens:", err)
		tx.Rollback()
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if err := tx.Where("id = ?", clientID).Delete(&models.ClientCredentials{}).Error; err != nil {
		log.Println("Error deleting client credentials:", err)
		tx.Rollback()
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Client successfully deleted"))
}
