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
	t, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	router.Get("/", controllers.StaticHandler(t))

	t, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	router.Get("/contact", controllers.StaticHandler(t))

	t, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	router.Get("/faq", controllers.StaticHandler(t))

	router.NotFoundHandler()
	fmt.Println("Starting server on 3000....")
	http.ListenAndServe(":3000", router)
}
