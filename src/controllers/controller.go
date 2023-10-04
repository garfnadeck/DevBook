package controllers

import (
	db2 "api/src/config/db"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(req, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db2.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.UserRepository(db)
	userID, err := repo.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userID)
	w.Write([]byte(fmt.Sprintf("ID created: %d", userID)))

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
