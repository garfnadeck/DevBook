package controllers

import (
	"api/src/config/db"
	"api/src/models"
	"api/src/repository"
	response_handler "api/src/response-handler"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"strings"
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

	if err = user.Prepare("register"); err != nil {
		response_handler.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repos := repository.UserRepository(db)
	user.ID, err = repos.Create(user)
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	response_handler.JSON(w, http.StatusCreated, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := db.ConnectDB()
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repos := repository.UserRepository(db)
	users, err := repos.SearchUser(nameOrNick)
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	response_handler.JSON(w, http.StatusOK, users)
}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repos := repository.UserRepository(db)
	user, err := repos.SearchById(userID)
	if err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	response_handler.JSON(w, http.StatusOK, user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		response_handler.Error(w, http.StatusBadRequest, err)
		return
	}
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
	if err = user.Prepare("editing"); err != nil {
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
	if err = repo.UpdateUser(userID, user); err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	response_handler.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
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
	if err = repo.DeleteUser(userID); err != nil {
		response_handler.Error(w, http.StatusInternalServerError, err)
		return
	}
	response_handler.JSON(w, http.StatusNoContent, nil)

}
