package handlers

import (
	"net/http"
	"time"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/serialization"
	"github.com/ortizdavid/go-rest-concepts/entities"
	"github.com/ortizdavid/go-rest-concepts/models"
)

type UserHandler struct {
}

func (us UserHandler) Routes(router *http.ServeMux) {
	router.HandleFunc("GET /api/users", us.getAllUsers)
	router.HandleFunc("GET /api/users/{id}", us.getUser)
	router.HandleFunc("POST /api/users", us.createUser)
	router.HandleFunc("PUT /api/users/{id}", us.updateUser)
	router.HandleFunc("DELETE /api/users/{id}", us.deleteUser)
}


func (UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var userModel models.UserModel
	users, err := userModel.FindAll()
	
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count := len(users)
	if count == 0 {
		httputils.WriteJsonError(w, "not found", http.StatusNotFound)
		return
	}
	httputils.WriteJson(w, http.StatusOK, users, count)
}


func (UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	var userModel models.UserModel
	id := r.PathValue("id")
	userId := conversion.StringToInt(id)

	user, err := userModel.FindById(userId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	httputils.WriteJsonSimple(w, http.StatusOK, user)
}


func (UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var userModel models.UserModel
	var user entities.User

	if err := serialization.DecodeJson(r.Body, &user); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	exists, _ := userModel.ExistsRecord("user_name", user.UserName)
	if exists {
		httputils.WriteJsonError(w, "Username'"+user.UserName+ "' exists", http.StatusConflict)
		return
	}

	user.UserId = userModel.LastInsertId
	user.Active = "Yes"
	user.Password = encryption.HashPassword(user.Password)
	user.UniqueId = encryption.GenerateUUID()
	user.Token = encryption.GenerateRandomToken()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := userModel.Create(user)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	httputils.WriteJsonSimple(w, http.StatusCreated, user)
}


func (UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
}


func (UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	
}