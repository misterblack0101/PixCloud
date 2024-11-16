package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"pixcloud/controllers"
	"pixcloud/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	t := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	router.Get("/", controllers.StaticHandler(t))

	t = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	router.Get("/contact", controllers.StaticHandler(t))

	t = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	router.Get("/faq", controllers.StaticHandler(t))

	fmt.Println("Starting server on 3000....")
	http.ListenAndServe(":3000", router)
}
