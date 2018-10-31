package routes

import (
	"encoding/json"
	"net/http"

	"rest-api/models"

	"github.com/gorilla/mux"
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

// ViewEmployee : handler function for GET /v1/employees/<id> call
func ViewEmployee(w http.ResponseWriter, r *http.Request) {
	accessToken := &models.AccessToken{}
	if !accessToken.AuthorizeByToken(w, r) {
		return
	}

	params := mux.Vars(r)
	employee := &models.Employee{}

	if employee.FindByID(w, params["id"]) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{"data": employee, "status": 1}
		json.NewEncoder(w).Encode(response)
	}

}

// DeleteEmployee : handler function for DELETE /v1/employees/<id> call
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
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
