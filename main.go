package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to my new website </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Contact page </h1>")
}

// type Router struct{}

// func (_ Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homePageHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.NotFound(w, r)
// 	}

// }

func main() {
	router := chi.NewRouter()
	router.Get("/", homePageHandler)
	router.Get("/contact", contactHandler)
	router.NotFoundHandler()
	fmt.Println("Starting server on 3000....")
	http.ListenAndServe(":3000", router)
}
