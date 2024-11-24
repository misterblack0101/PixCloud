package controllers

import (
	"fmt"
	"net/http"
	"pixcloud/models"
)

type Users struct {
	// anonymous struct called Tempates
	Templates struct {
		New Template
	}
	UserService *models.UserService
}

func (user Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	user.Templates.New.Execute(w, data)
}

func (user Users) Create(w http.ResponseWriter, r *http.Request) {
	email, password := r.PostFormValue("email"), r.PostFormValue("password")
	// create new user
	userModel, err := user.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong while creating the user!", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "User Created successfully", userModel)

}
