package main

import (
	"fmt"
	"net/http"
	"pixcloud/controllers"
	"pixcloud/models"
	"pixcloud/templates"
	"pixcloud/views"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()
	t := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	router.Get("/", controllers.StaticHandler(t))

	t = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	router.Get("/contact", controllers.StaticHandler(t))

	t = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	router.Get("/faq", controllers.FAQ(t))

	// create connection to db
	cfg := models.DefaultPostgresConfig()
	db, err := models.OpenDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	userStruct := controllers.Users{
		UserService: &userService,
	}
	userStruct.Templates.New =
		views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))

	router.Get("/signup", userStruct.New)
	router.Post("/users", userStruct.Create)

	router.NotFound(controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "not_found.gohtml"))))

	fmt.Println("Starting server on 3000....")
	http.ListenAndServe(":3000", router)
}
