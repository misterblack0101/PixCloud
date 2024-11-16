package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTemplate *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(filePath string) (Template, error) {
	tmp, err := template.ParseFiles(filePath)
	if err != nil {
		return Template{}, fmt.Errorf("error while parsing template %w", err)
	}
	return Template{htmlTemplate: tmp}, nil

}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	err := t.htmlTemplate.Execute(w, data)
	if err != nil {
		log.Print("error parsing template", err)
		http.Error(w, "<h1> There was en error loading the site</h1>", http.StatusInternalServerError)
		return
	}
}
