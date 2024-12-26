package v1_clients

import (
	"log"
	"net/http"

	"backend/middleware"
	"backend/utils"
)

func LogoutClientHandler(w http.ResponseWriter, r *http.Request) {
	clientID := r.Context().Value(middleware.ClientIDKey).(int)

	db := utils.GetDB()

	_, err := db.Exec("DELETE FROM client_refresh_tokens WHERE client_id = $1", clientID)
	if err != nil {
		log.Println("Error removing refresh tokens:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
