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

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string        `bson:"name" json:"name"`
	Username  string        `bson:"username" json:"username"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"password,omitempty"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

func (user *User) Validate(w http.ResponseWriter, r *http.Request) bool {
	errs := url.Values{}
	db := database.Db

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errs.Add("data", "Invalid data")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"errors": errs, "status": 0}
		json.NewEncoder(w).Encode(response)
		return false
	}

	if govalidator.IsNull(user.Name) {
		errs.Add("name", "Name is required")
	}
	if govalidator.IsNull(user.Username) {
		errs.Add("username", "Username is required")
	}
	if govalidator.IsNull(user.Email) {
		errs.Add("email", "E-mail is required")
	}
	if govalidator.IsNull(user.Password) {
		errs.Add("password", "Password is required")
	}

	count, _ := db.C("user").Find(bson.M{"email": user.Email}).Count()
	if count > 0 {
		errs.Add("email", "E-mail is already in use")
	}
	count, _ = db.C("user").Find(bson.M{"username": user.Username}).Count()
	if count > 0 {
		errs.Add("username", "Username is already in use")
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
