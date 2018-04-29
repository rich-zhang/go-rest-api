package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var users []User

func populateData() {
	users = append(users, User{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	users = append(users, User{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	users = append(users, User{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func GetUser(w http.ResponseWriter, r *http.Request)    {}
func GetUsers(w http.ResponseWriter, r *http.Request)   {}
func CreateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}

// our main function
func main() {
	// populate users with stub data
	populateData()

	router := mux.NewRouter()
	router.HandleFunc("/users", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
