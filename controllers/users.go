package controllers

import (
	"fmt"
	"net/http"
	"pixcloud/models"
)

type Users struct {
	// anonymous struct called Tempates
	Templates struct {
		New    Template
		SignIn Template
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

func (user Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	user.Templates.SignIn.Execute(w, data)
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

func (user Users) Login(w http.ResponseWriter, r *http.Request) {
	email, password := r.PostFormValue("email"), r.PostFormValue("password")
	fmt.Println(email, password)
	// create new user
	userModel, err := user.UserService.Login(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid credentials!", http.StatusUnauthorized)
		return
	}
	cookie := http.Cookie{Name: "email", Value: userModel.Email, Path: "/"}
	http.SetCookie(w, &cookie)
	fmt.Fprint(w, "User logged in successfully", userModel)
}

func (user Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")
	if err != nil {
		fmt.Fprint(w, "Cannot find user cookie")
		return
	}
	fmt.Fprint(w, "User email is %s", email.Value)
}
