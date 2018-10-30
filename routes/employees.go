package routes

import (
	"encoding/json"
	"net/http"

	"rest-api/models"
)

/*
func Register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if !user.Validate(w, r) {
		return
	}
	db := database.Db
	c := db.C("user")
	// Insert
	user.ID = bson.NewObjectId()
	salt, _ := bcrypt.Salt(10)
	user.Password, _ = bcrypt.Hash(user.Password, salt)

	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()

	insertionErrors := c.Insert(&user)

	if insertionErrors != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"errors": insertionErrors.Error(), "status": 0}
		json.NewEncoder(w).Encode(response)

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{"data": user, "status": 1}
		json.NewEncoder(w).Encode(response)
	}

}

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
*/

// CreateEmployee : handler function for POST /v1/employees call
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	accessToken := &models.AccessToken{}
	if !accessToken.AuthorizeByToken(w, r) {
		return
	}

	employee := &models.Employee{}

	if employee.Validate(w, r) && employee.Create(w) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{"data": employee, "status": 1}
		json.NewEncoder(w).Encode(response)
	}

}
