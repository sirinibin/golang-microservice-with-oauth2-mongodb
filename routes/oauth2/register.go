package register

func (user *User) validate(w http.ResponseWriter,r *http.Request) bool  {
	errs := url.Values{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	   errs.Add("data", "Invalid data")

	   w.WriteHeader(http.StatusBadRequest)
	   response := map[string]interface{}{"errors":errs,"status": 0}
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

   if len(errs)>0 {
	   w.WriteHeader(http.StatusBadRequest)
	   response := map[string]interface{}{"errors":errs,"status": 0}
	   json.NewEncoder(w).Encode(response)
		return false
   }
   return true
}