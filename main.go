package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	tmp, err := template.ParseFiles(filePath)
	if err != nil {
		log.Print("Error parsing template\n", err)
		http.Error(w, "<h1> There was en error loading the site</h1>", http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		log.Print("Error parsing template\n", err)
		http.Error(w, "<h1> There was en error loading the site</h1>", http.StatusInternalServerError)
		return
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "home.gohtml"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "contact.gohtml"))
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func main() {
	router := chi.NewRouter()
	router.Get("/", homePageHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.NotFoundHandler()
	fmt.Println("Starting server on 3000....")
	http.ListenAndServe(":3000", router)
}
