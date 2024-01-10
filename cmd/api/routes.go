package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Post("/books", app.postBookHandler)
		r.Get("/books/{id}", app.getBookHandler)
		r.Put("/books/{id}", app.putBookHandler)
		r.Delete("/books/{id}", app.deleteBookHandler)
		r.Get("/books", app.getAllBooksHandler)
	})

	return router
}
