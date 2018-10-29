package routes

import (
	"encoding/json"
	"net/http"
	"rest-api/models"
)

// LogOut : handler function for /v1/logout call
func LogOut(w http.ResponseWriter, r *http.Request) {
	accessToken := &models.AccessToken{}
	if !accessToken.AuthorizeByToken(w, r) {
		return
	}
	user := accessToken.GetUser()
	accessToken.Remove()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"data": map[string]interface{}{"user_id": user.ID, "message": "LoggedOut Successfully"}, "status": 1}
	json.NewEncoder(w).Encode(response)
}
