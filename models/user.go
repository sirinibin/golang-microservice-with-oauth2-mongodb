package user

import (
	"encoding/json"	
	"gopkg.in/mgo.v2/bson"
   )

   
type User struct {
	ID  bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
	Username string `bson:"username" json:"username"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}