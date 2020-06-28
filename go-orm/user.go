package main


import(
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "All Users End point hit!")
}

func NewUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "New User End point hit!")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Delete User End point hit!")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Update User End point hit!")
}

