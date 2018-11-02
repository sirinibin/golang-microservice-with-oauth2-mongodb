package routes

import (
	"encoding/json"
	"net/http"

	"rest-api/models"
)

// Authorize : handler function for /v1/authorize api call
func Authorize(w http.ResponseWriter, r *http.Request) {

	authRequest := &models.AuthorizeRequestBody{}
	if !authRequest.Validate(w, r) {
		return
	}
	authcode := authRequest.GenerateAuthCode(w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{"data": map[string]interface{}{
		"authorization_code": authcode.Code,
		"expires_at":         authcode.ExpiresAt,
	}, "status": 1}
	json.NewEncoder(w).Encode(response)
}

// AccessToken : handler function for /v1/accesstoken call
func AccessToken(w http.ResponseWriter, r *http.Request) {

	accessTokenRequest := &models.AccessTokenRequestBody{}
	if !accessTokenRequest.Validate(w, r) {
		return
	}
	accesstoken := accessTokenRequest.GenerateAccessToken(w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{"data": map[string]interface{}{
		"access_token": accesstoken.Token,
		"expires_at":   accesstoken.ExpiresAt,
	}, "status": 1}

	json.NewEncoder(w).Encode(response)
}
