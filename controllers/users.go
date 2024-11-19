package controllers

import (
	"net/http"
	"pixcloud/views"
)

type Users struct {
	// anonymous struct called Tempates
	Templates struct {
		New views.Template
	}
}

func (user Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	user.Templates.New.Execute(w, nil)
}
