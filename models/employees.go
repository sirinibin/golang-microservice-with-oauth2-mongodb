package models

import (
	"encoding/json"
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
	CreatedAt time.Time     `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

var err error

//Remove : remove access token
func (employee *Employee) Remove(w http.ResponseWriter) bool {

	db := database.Db
	err := db.C("employees").Remove(&employee)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"errors": err.Error(), "status": 0}
		json.NewEncoder(w).Encode(response)
		return false
	}
	return true

}

// FindByID : Find Employee record
func (employee *Employee) FindByID(w http.ResponseWriter, id string) bool {

	errs := url.Values{}
	db := database.Db

	if bson.IsObjectIdHex(id) {
		err = db.C("employees").FindId(bson.ObjectIdHex(id)).One(&employee)
		if err != nil {
			errs.Add("id", "Invalid Document ID")
		}

	} else {
		errs.Add("id", "Invalid Document ID")
	}

	//fmt.Printf("%+v\n", id)
	if len(errs) > 0 {
		//errs.Add("id", "Invalid Document ID")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"errors": errs, "status": 0}
		json.NewEncoder(w).Encode(response)
		return false
	}

	return true

}

// Save : Create/Update Employee record
func (employee *Employee) Save(w http.ResponseWriter) bool {

	db := database.Db
	c := db.C("employees")

	employee.UpdatedAt = time.Now().Local()

	if employee.ID == "" {
		employee.ID = bson.NewObjectId()
		employee.CreatedAt = time.Now().Local()
		err = c.Insert(&employee)

	} else {

		err = c.Update(bson.M{"_id": employee.ID}, bson.M{"$set": employee})
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"errors": err.Error(), "status": 0}
		json.NewEncoder(w).Encode(response)
		return false

	} else {

		return true
	}
}

// Validate : Validate Employee creation/Updation
func (employee *Employee) Validate(w http.ResponseWriter, r *http.Request, action string) bool {
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

	if action == "update" && employee.ID == "" {
		errs.Add("id", "id is required")
	}
	if action == "update" && employee.ID != "" {
		oldEmployee := Employee{}
		err = db.C("employees").Find(bson.M{"_id": employee.ID}).One(&oldEmployee)
		employee.CreatedAt = oldEmployee.CreatedAt
		if err != nil {
			errs.Add("id", "Invalid Document ID")
		}

	}

	if govalidator.IsNull(employee.Name) {
		errs.Add("name", "Name is required")
	}
	if govalidator.IsNull(employee.Email) {
		errs.Add("email", "E-mail is required")
	}

	count := 0
	if action == "create" {
		//New Record
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
