package main

import (
	"fmt"
	"github.com/xuche123/bookwise/internal/data"
	"net/http"
	"time"
)

func (app *application) postBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDFromParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	book := data.Book{
		ID:            id,
		Title:         "Test book",
		Author:        "Test author",
		PublishedDate: time.Time{},
		ImageURL:      "-",
		Description:   "Test description",
		CreatedAt:     time.Now(),
		Version:       1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
