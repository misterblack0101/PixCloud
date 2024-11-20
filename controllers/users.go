package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	// anonymous struct called Tempates
	Templates struct {
		New Template
	}
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
	fmt.Fprint(w, email, password)

}
