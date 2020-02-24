package main

import (
	"encoding/json"
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
	users = append(users, User{
		ID:        "1",
		Firstname: "Frank",
		Lastname:  "Willis",
		Address:   &Address{City: "A City", State: "B State"},
	})
	users = append(users, User{
		ID:        "2",
		Firstname: "Michelle",
		Lastname:  "Doe",
		Address: &Address{
			City:  "B City",
			State: "B State"},
	})
	users = append(users, User{
		ID:        "3",
		Firstname: "William",
		Lastname:  "Sully"},
	)
}

func GetUser(w http.ResponseWriter, r *http.Request)    {}
func GetUsers(w http.ResponseWriter, r *http.Request)   {}
func CreateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}

// get all users
// return json
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

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
