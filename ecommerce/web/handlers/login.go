package handlers

import (
	"ecommerce/auth"
	"ecommerce/db"
	model "ecommerce/models"
	"ecommerce/web/utils"
	"encoding/json"
	"net/http"
)

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user LoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = validate(user)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err, "Validation error")
		return
	}

	err = db.Login(user.Email, user.Password)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err, "Wrong username / password ")
		return
	}

	usrchan := make(chan model.User)   //channel
	go db.GetUser(user.Email, usrchan) //goroutine
	usr := <-usrchan                   // get user from goroutine

	token, err := auth.GenerateToken(usr)

	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	utils.SendBothData(w, token, usr)
}
