package main

import (
	"fmt"
	"github.com/xuche123/bookwise/internal/data"
	"github.com/xuche123/bookwise/internal/validator"
	"net/http"
	"time"
)

func (app *application) postBookHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		ImageURL    string `json:"image_url"`
		Description string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	book := &data.Book{
		Title:       input.Title,
		Author:      input.Author,
		ImageURL:    input.ImageURL,
		Description: input.Description,
	}

	v := validator.New()

	if data.ValidateBook(v, book); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	_, err = fmt.Fprintf(w, "%+v\n", input)
	if err != nil {
		return
	}
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDFromParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	book := data.Book{
		ID:          id,
		Title:       "Test book",
		Author:      "Test author",
		ImageURL:    "-",
		Description: "Test description",
		CreatedAt:   time.Now(),
		Version:     1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
