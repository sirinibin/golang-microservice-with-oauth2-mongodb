package routes

import (
	"encoding/json"
	"net/http"
	"rest-api/db"

	"gopkg.in/mgo.v2/bson"

	"rest-api/models"

	"github.com/jameskeane/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	insertionErrors := c.Insert(&user)

	if insertionErrors != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"errors": insertionErrors.Error(), "status": 0}
		json.NewEncoder(w).Encode(response)

	} else {
		w.WriteHeader(http.StatusOK)
		user.Password = ""
		response := map[string]interface{}{"data": user, "status": 1}
		json.NewEncoder(w).Encode(response)
	}

}
