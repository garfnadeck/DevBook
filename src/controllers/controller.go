package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get All User"))
}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Single User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}
