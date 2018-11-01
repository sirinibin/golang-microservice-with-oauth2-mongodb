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
	router.HandleFunc("/v1/employees", routes.ListEmployees).Methods("GET")
	router.HandleFunc("/v1/employees/{id}", routes.ViewEmployee).Methods("GET")
	router.HandleFunc("/v1/employees", routes.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/v1/employees/{id}", routes.DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8008", router))
}
