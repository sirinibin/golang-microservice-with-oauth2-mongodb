package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"rest-api/db"
	time "time"

	"github.com/asaskevich/govalidator"
	"gopkg.in/mgo.v2/bson"
)

// Employee : struct for Employee model
type Employee struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

// Create : Create Employee record
func (employee *Employee) Create(w http.ResponseWriter) bool {

	db := database.Db
	c := db.C("employees")

	employee.ID = bson.NewObjectId()
	employee.CreatedAt = time.Now().Local()
	employee.UpdatedAt = time.Now().Local()

	insertionErrors := c.Insert(&employee)

	if insertionErrors != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"errors": insertionErrors.Error(), "status": 0}
		json.NewEncoder(w).Encode(response)
		return false

	} else {

		return true
	}
}

// Validate : Validate Employee creation/Updation
func (employee *Employee) Validate(w http.ResponseWriter, r *http.Request) bool {
	errs := url.Values{}
	db := database.Db

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		errs.Add("data", "Invalid data")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"errors": errs, "status": 0}
		json.NewEncoder(w).Encode(response)
		return false
	}

	if govalidator.IsNull(employee.Name) {
		errs.Add("name", "Name is required")
	}
	if govalidator.IsNull(employee.Email) {
		errs.Add("email", "E-mail is required")
	}

	fmt.Printf("%+v\n", employee)

	count := 0
	if employee.ID == "" {
		//New Records
		count, _ = db.C("employees").Find(bson.M{"email": employee.Email}).Count()
	} else {
		//Existing Record
		count, _ = db.C("employees").Find(bson.M{"email": employee.Email, "_id": bson.M{"$ne": employee.ID}}).Count()
	}

	if count > 0 {
		errs.Add("email", "E-mail is already in use")
	}
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"errors": errs, "status": 0}
		json.NewEncoder(w).Encode(response)
		return false
	}
	return true
}
