package main

import (
 "encoding/json"
 "log"
 "net/http"
 "github.com/gorilla/mux"
 mgo "gopkg.in/mgo.v2"
 "gopkg.in/mgo.v2/bson"
 "github.com/asaskevich/govalidator"
 //"errors"
 "net/url"
 "routes/oauth2/register"
 //"fmt"

)

var db *mgo.Database

var employees []Employee

func main() {

 ConnectToDb()
 router :=	mux.NewRouter()
 router.HandleFunc("/v1/register",Register).Methods("POST")
 router.HandleFunc("/v1/employees",GetEmployees).Methods("GET")
 router.HandleFunc("/v1/employees/{id}",GetEmployee).Methods("GET")
 router.HandleFunc("/v1/employees",CreateEmployee).Methods("POST")
 router.HandleFunc("/v1/employees/{id}",UpdateEmployee).Methods("PUT")
 router.HandleFunc("/v1/employees/{id}",DeleteEmployee).Methods("DELETE")


 //employees = append(employees, Employee{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
 //employees = append(employees, Employee{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
 //employees = append(employees, Employee{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
 

 log.Fatal(http.ListenAndServe(":8008",router))
}


func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	user := &User{}
	if !user.validate(w,r)  {
		return
	}
	  
	c := db.C("user")
	// Insert
	user.ID = bson.NewObjectId()
	insertionErrors := c.Insert(&user)
	if insertionErrors != nil {
		//panic(err)
	
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"errors":insertionErrors.Error(),"status": 0}
		json.NewEncoder(w).Encode(response)

	}else {
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{"data":user,"status": 1}
		json.NewEncoder(w).Encode(response)
     }
	
   
   }

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","application/json")
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
  */
  json.NewEncoder(w).Encode(&Employee{})

}
func CreateEmployee(w http.ResponseWriter, r *http.Request) {

 var employee Employee
 _ = json.NewDecoder(r.Body).Decode(&employee)


 c := db.C("employees")
 // Insert
 employee.ID = bson.NewObjectId()
 err := c.Insert(&employee)
 if err != nil {
	 panic(err)
 }

 w.Header().Set("Content-Type","application/json")
 w.WriteHeader(http.StatusOK)
 json.NewEncoder(w).Encode(employee)

}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)

	var employee Employee
    _ = json.NewDecoder(r.Body).Decode(&employee)

	c := db.C("employees")
	// Update
    //employee.ID = bson.NewObjectId()
	//err := c.Insert(&employee)
	err := c.UpdateId(employee.ID, &employee)

	if err != nil {
		//panic(err)
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"errors":err.Error(),"status": 0}
		json.NewEncoder(w).Encode(response)

	}else {

		w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"data":employee,"status": 1}
	json.NewEncoder(w).Encode(response)
	}
   
	

}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)
	/*
    for index, item := range employees {
        if item.ID == params["id"] {
            employees = append(employees[:index], employees[index+1:]...)
            break
        }
    
	} */
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
	return
}
/*
type User struct {
	ID  bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name" valid:"required~Name is required" `
	Username string `bson:"username" json:"username" valid:"required~Username is required"`
	Email string `bson:"email" json:"email" valid:"required~E-mail is required"`
	Password string `bson:"password" json:"password" valid:"required~Password is required"`
} */
type User struct {
	ID  bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
	Username string `bson:"username" json:"username"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type Employee struct {
	ID  bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}
type EmployeesDAO struct {
	Server   string
	Database string
}
type Response struct {
	status   int
	data  Employee `json:"data"`
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




func ConnectToDb() {

	Host := []string{
		"localhost:27017",
		// replica set addrs...
	}
	const (
		Database   = "expressjs_api"
		Collection = "employees"
	)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
		// Username: Username,
		// Password: Password,
		// Database: Database,
		// DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
		// 	return tls.Dial("tcp", addr.String(), &tls.Config{})
		// },
	})
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	db = session.DB(Database)

}