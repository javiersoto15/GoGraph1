package router

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/temp/GQL-Tut/api/handler"
)

//Initiliaze will handle the routes for the chi server
func initiliaze() *chi.Mux {
	router := chi.NewRouter()

	//uses chi middleware logic
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	//Sets context for all requests
	router.Use(middleware.Timeout(20 * time.Second))

	// "/v1" groups queries for version v1
	// "/" handlers is for graphiQL
	// "/query" is for queries
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", handler.Routes())
	})
	return router
}

//ServeRouter initiliaze the router and serves
func ServeRouter() {
	r := initiliaze()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Error serving")
	}
}
