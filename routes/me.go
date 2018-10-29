package routes

import (
	"encoding/json"
	"net/http"
	"rest-api/models"
)

// Me : handler function for /v1/me call
func Me(w http.ResponseWriter, r *http.Request) {
	accessToken := &models.AccessToken{}
	if !accessToken.AuthorizeByToken(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	user := accessToken.GetUser()
	response := map[string]interface{}{"data": user, "status": 1}
	json.NewEncoder(w).Encode(response)
}
