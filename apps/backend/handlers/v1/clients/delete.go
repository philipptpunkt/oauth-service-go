package v1_clients

import (
	"log"
	"net/http"

	"backend/middleware"
	"backend/utils"
)

func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientID := r.Context().Value(middleware.ClientIDKey).(int)

	db := utils.GetDB()
	_, err := db.Exec("DELETE FROM client_credentials WHERE id = $1", clientID)
	if err != nil {
		log.Println("Error deleting client:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM client_refresh_tokens WHERE client_id = $1", clientID)
	if err != nil {
		log.Println("Error deleting refresh tokens:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Client successfully deleted"))
}
