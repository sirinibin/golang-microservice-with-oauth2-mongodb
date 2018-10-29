package routes

import (
	"encoding/json"
	"net/http"
	"rest-api/db"
	"time"

	"gopkg.in/mgo.v2/bson"

	"rest-api/models"

	"github.com/jameskeane/bcrypt"
)

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
