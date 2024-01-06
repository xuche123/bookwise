package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Post("/books", app.postBookHandler)
		r.Get("/books/{id}", app.getBookHandler)
	})

	return router
}
