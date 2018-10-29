package routes

import (
	"encoding/json"
	"net/http"

	"rest-api/models"
)

// Authorize : handler function for /v1/authorize api call
func Authorize(w http.ResponseWriter, r *http.Request) {

	auth := &models.Authorize{}
	if !auth.Validate(w, r) {
		return
	}
	authcode := auth.GenerateAuthCode(w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{"data": map[string]interface{}{
		"authorization_code": authcode.Code,
		"expires_at":         authcode.ExpiresAt,
	}, "status": 1}
	json.NewEncoder(w).Encode(response)
}
