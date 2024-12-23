package v1

import (
	"encoding/json"
	"log"
	"net/http"
)

type VerifyCodeRequest struct{}

type VerifyCodeResponse struct {
	Message string `json:"message"`
}

func TestCookieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookies := r.Cookies()
	log.Println("Read Cookie:")
	for _, c := range cookies {
		log.Printf("Cookie: %s=%s\n", c.Name, c.Value)
	}

	_, err := r.Cookie("test_cookie")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VerifyCodeResponse{
			Message: "NOT OK",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VerifyCodeResponse{
		Message: "OK",
	})
}
