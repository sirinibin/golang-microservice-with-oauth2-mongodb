package main

import (
	"log"
	"net/http"
	"rest-api/db"
	"rest-api/routes"

	"github.com/gorilla/mux"
)

//var employees []Employee

func main() {

	database.Connect()
	router := mux.NewRouter()

	//Register
	router.HandleFunc("/v1/register", routes.Register).Methods("POST")

	//OAuth2
	router.HandleFunc("/v1/authorize", routes.Authorize).Methods("POST")
	router.HandleFunc("/v1/accesstoken", routes.AccessToken).Methods("POST")

	//User
	router.HandleFunc("/v1/me", routes.Me).Methods("GET")
	router.HandleFunc("/v1/logout", routes.LogOut).Methods("GET")

	//Employees
	router.HandleFunc("/v1/employees", routes.CreateEmployee).Methods("POST")
	router.HandleFunc("/v1/employees", routes.UpdateEmployee).Methods("PUT")

	/*
		router.HandleFunc("/v1/employees", GetEmployees).Methods("GET")
		router.HandleFunc("/v1/employees/{id}", GetEmployee).Methods("GET")
		router.HandleFunc("/v1/employees", CreateEmployee).Methods("POST")
		router.HandleFunc("/v1/employees/{id}", UpdateEmployee).Methods("PUT")
		router.HandleFunc("/v1/employees/{id}", DeleteEmployee).Methods("DELETE")
	*/

	log.Fatal(http.ListenAndServe(":8008", router))
}

/*
func GetEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write(employee)
	json.NewEncoder(w).Encode(employees)
}
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	/*
		 for _,item := range employees {
			 if item.ID == params["id"] {
				w.Header().Set("Content-Type","application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(item)
				return
			 }
		  }

	json.NewEncoder(w).Encode(&Employee{})

}
func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	/*
		var employee Employee
		_ = json.NewDecoder(r.Body).Decode(&employee)

		c := db.C("employees")
		// Insert
		employee.ID = bson.NewObjectId()
		err := c.Insert(&employee)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employee)


}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)

	/*
		var employee Employee
		_ = json.NewDecoder(r.Body).Decode(&employee)

		c := db.C("employees")
		// Update
		//employee.ID = bson.NewObjectId()
		//err := c.Insert(&employee)
		err := c.UpdateId(employee.ID, &employee)

		if err != nil {
			//panic(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]interface{}{"errors": err.Error(), "status": 0}
			json.NewEncoder(w).Encode(response)

		} else {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			response := map[string]interface{}{"data": employee, "status": 1}
			json.NewEncoder(w).Encode(response)
		}
	*

}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)
	/*
	    for index, item := range employees {
	        if item.ID == params["id"] {
	            employees = append(employees[:index], employees[index+1:]...)
	            break
	        }

		}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
	return
}


type Employee struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string        `bson:"name" json:"name"`
	Email string        `bson:"email" json:"email"`
}
type EmployeesDAO struct {
	Server   string
	Database string
}
type Response struct {
	status int
	data   Employee `json:"data"`
}

//employees = append(employees, Employee{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
/*
 var db *mgo.Database
const (
	COLLECTION = "employees"
	)
func (m *EmployeesDAO) Connect() {

	session, err := mgo.Dial(m.Server)
	if err != nil {
	log.Fatal(err)
    }
	db = session.DB(m.Database)
}
*/
