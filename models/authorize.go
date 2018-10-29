package models

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"net/url"
	"rest-api/db"
	time "time"

	"github.com/asaskevich/govalidator"
	"github.com/jameskeane/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var user User

//Authorize : Authorize structure
type Authorize struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

//Authcode : Authcode structure
type Authcode struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Code      []byte        `bson:"code" json:"code"`
	ExpiresAt time.Time     `bson:"expires_at" json:"expires_at"`
	UserID    bson.ObjectId `bson:"user_id" json:"user_id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

//GetUser : return user object
func (auth *Authorize) GetUser() User {
	return user
}

//GenerateAuthCode : generate and return authcode
func (auth *Authorize) GenerateAuthCode(w http.ResponseWriter) *Authcode {
	b := make([]byte, 50)
	rand.Read(b)

	authcode := &Authcode{}
	authcode.ID = bson.NewObjectId()
	authcode.Code = b
	authcode.UserID = user.ID
	authcode.ExpiresAt = time.Now().Local().Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(5) +
		time.Second*time.Duration(0))
	authcode.CreatedAt = time.Now().Local()
	authcode.UpdatedAt = time.Now().Local()

	db := database.Db
	c := db.C("authcode")
	// Insert
	insertionErrors := c.Insert(&authcode)

	if insertionErrors != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"errors": insertionErrors.Error(), "status": 0}
		json.NewEncoder(w).Encode(response)

	}

	return authcode
}

//Validate : Validate authorization data
func (auth *Authorize) Validate(w http.ResponseWriter, r *http.Request) bool {
	errs := url.Values{}
	db := database.Db

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		errs.Add("data", "Invalid data")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"errors": errs, "status": 0}
		json.NewEncoder(w).Encode(response)
		return false
	}

	if govalidator.IsNull(auth.Username) {
		errs.Add("username", "Username is required")
	}
	if govalidator.IsNull(auth.Password) {
		errs.Add("password", "Password is required")
	}

	if !govalidator.IsNull(auth.Password) && !govalidator.IsNull(auth.Username) {

		err := db.C("user").Find(bson.M{"username": auth.Username}).One(&user)
		if err != nil {
			errs.Add("password", "Username or Password is wrong")
		} else {
			if !bcrypt.Match(auth.Password, user.Password) {
				errs.Add("password", "Username or Password is wrong")
			}
		}
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
