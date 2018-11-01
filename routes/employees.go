package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// ListEmployees : handler function for GET /v1/employees call
func ListEmployees(w http.ResponseWriter, r *http.Request) {
	accessToken := &models.AccessToken{}
	if !accessToken.AuthorizeByToken(w, r) {
		return
	}
	//params := mux.Vars(r)

	employee := &models.Employee{}

	limit := 10
	page := 1
	order := "name"
	//order := bson.D{{Name: "name", Value: -1}}

	keys, ok := r.URL.Query()["limit"]
	if ok && len(keys[0]) >= 1 {
		limit, _ = strconv.Atoi(keys[0])
	}
	keys, ok = r.URL.Query()["page"]
	if ok && len(keys[0]) >= 1 {
		page, _ = strconv.Atoi(keys[0])
	}
	keys, ok = r.URL.Query()["order"]
	if ok && len(keys[0]) >= 1 {
		order = keys[0]
	}

	//fmt.Printf("%v", keys)
	//fmt.Printf("%+v\n", keys[0])

	offset := (page - 1) * limit

	if employee.FindAll(w, offset, limit, order) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{"data": employee.GetEmployeeList(), "status": 1}
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
	params := mux.Vars(r)
	employee := &models.Employee{}

	if employee.FindByID(w, params["id"]) && employee.Remove(w) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{"data": employee, "message": "Deleted Successfully", "status": 1}
		json.NewEncoder(w).Encode(response)
	}

}
