package models

import (
	"encoding/json"
	"net/http"
	"net/url"
	"rest-api/db"
	time "time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

//var accessToken AccessToken
var authocode Authcode

//AccessTokenRequestBody : AccessTokenRequestBody structure
type AccessTokenRequestBody struct {
	AuthorizationCode string `bson:"authorization_code" json:"authorization_code"`
}

//AccessToken : Authcode structure
type AccessToken struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Token     string        `bson:"code" json:"code"`
	ExpiresAt time.Time     `bson:"expires_at" json:"expires_at"`
	UserID    bson.ObjectId `bson:"user_id" json:"user_id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

//GetAccessToken : return GetAccessToken object
/*
func (accessTokenRequest *AccessTokenRequestBody) GetAccessToken() AccessToken {
	return accessToken
}
*/

//GenerateAccessToken : generate and return accessToken Object
func (accessTokenRequest *AccessTokenRequestBody) GenerateAccessToken(w http.ResponseWriter) *AccessToken {
	accessToken := &AccessToken{}
	accessToken.ID = bson.NewObjectId()
	accessToken.Token = uuid.Must(uuid.NewV4()).String()
	accessToken.UserID = user.ID
	accessToken.ExpiresAt = time.Now().Local().Add(time.Hour*time.Duration((24*60)) +
		time.Minute*time.Duration(0) +
		time.Second*time.Duration(0))
	accessToken.CreatedAt = time.Now().Local()
	accessToken.UpdatedAt = time.Now().Local()

	db := database.Db
	c := db.C("accesstoken")
	// Insert
	insertionErrors := c.Insert(&accessToken)

	if insertionErrors != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"errors": insertionErrors.Error(), "status": 0}
		json.NewEncoder(w).Encode(response)

	}

	return accessToken
}

//Validate : Validate authorization data
func (accessTokenRequest *AccessTokenRequestBody) Validate(w http.ResponseWriter, r *http.Request) bool {
	errs := url.Values{}
	db := database.Db

	if err := json.NewDecoder(r.Body).Decode(&accessTokenRequest); err != nil {
		errs.Add("data", "Invalid data")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"errors": errs, "status": 0}
		json.NewEncoder(w).Encode(response)
		return false
	}
	if govalidator.IsNull(accessTokenRequest.AuthorizationCode) {
		errs.Add("authorization_code", "Authorization Code is required")
	}

	if !govalidator.IsNull(accessTokenRequest.AuthorizationCode) {

		now := time.Now().Local()
		err := db.C("authcode").Find(bson.M{"code": accessTokenRequest.AuthorizationCode, "expires_at": bson.M{"$gt": now}}).One(&authocode)
		if err != nil {
			errs.Add("authorization_code", "Invalid Auth code")
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
