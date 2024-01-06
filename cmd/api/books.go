package main

import (
	"fmt"
	"net/http"
)

func (app *application) postBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDFromParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of book %d\n", id)
}
