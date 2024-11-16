package main

import (
	"fmt"
	"net/http"
	"pixcloud/controllers"
	"pixcloud/templates"
	"pixcloud/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	t := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	router.Get("/", controllers.StaticHandler(t))

	t = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	router.Get("/contact", controllers.StaticHandler(t))

	t = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	router.Get("/faq", controllers.StaticHandler(t))

	fmt.Println("Starting server on 3000....")
	http.ListenAndServe(":3000", router)
}
