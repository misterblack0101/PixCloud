package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"pixcloud/views"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	t, err := views.Parse(filePath)
	if err != nil {
		log.Print(err)
		http.Error(w, "<h1> There was en error loading the site</h1>", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
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
