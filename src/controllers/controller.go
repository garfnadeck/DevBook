package controllers

import (
	"api/src/config/db"
	"api/src/models"
	"api/src/repository"
	response_handler "api/src/response-handler"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		response_handler.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(req, &user); err != nil {
		response_handler.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		response_handler.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.UserRepository(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	response_handler.JSON(w, http.StatusCreated, user)
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
