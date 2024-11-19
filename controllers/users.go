package controllers

import (
	"net/http"
)

type Users struct {
	// anonymous struct called Tempates
	Templates struct {
		New Template
	}
}

func (user Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	user.Templates.New.Execute(w, nil)
}
