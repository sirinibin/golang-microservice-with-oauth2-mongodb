package routes

import (
	"encoding/json"
	"net/http"

	"rest-api/models"
)

// CreateEmployee : handler function for POST /v1/employees call
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	accessToken := &models.AccessToken{}
	if !accessToken.AuthorizeByToken(w, r) {
		return
	}

	employee := &models.Employee{}

	if employee.Validate(w, r, "create") && employee.Save(w) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{"data": employee, "status": 1}
		json.NewEncoder(w).Encode(response)
	}

}

// UpdateEmployee : handler function for PUT /v1/employees call
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	accessToken := &models.AccessToken{}
	if !accessToken.AuthorizeByToken(w, r) {
		return
	}

	employee := &models.Employee{}

	if employee.Validate(w, r, "update") && employee.Save(w) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{"data": employee, "status": 1}
		json.NewEncoder(w).Encode(response)
	}

}
