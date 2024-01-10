package main

import (
	"errors"
	"fmt"
	"github.com/xuche123/bookwise/internal/data"
	"github.com/xuche123/bookwise/internal/validator"
	"net/http"
	"strconv"
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

	err = app.models.Books.Insert(book)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/book/%d", book.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"book": book}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDFromParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	book, err := app.models.Books.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) putBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDFromParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	book, err := app.models.Books.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if r.Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(book.Version), 32) != r.Header.Get("X-Expected-Version") {
			app.editConflictResponse(w, r)
			return
		}
	}

	var input struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		ImageURL    string `json:"image_url"`
		Description string `json:"description"`
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	book.Title = input.Title
	book.Author = input.Author
	book.ImageURL = input.ImageURL
	book.Description = input.Description

	v := validator.New()

	if data.ValidateBook(v, book); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Books.Update(book)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDFromParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Books.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "book successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string
		Author string
		data.Filter
	}

	params := r.URL.Query()

	v := validator.New()
	input.Title = app.readString(params, "title", "")
	input.Author = app.readString(params, "author", "")
	input.Filter.Page = app.readInt(params, "page", 1, v)
	input.Filter.PageSize = app.readInt(params, "page_size", 20, v)
	input.Filter.Sort = app.readString(params, "sort", "id")
	input.Filter.SortSafeList = []string{"id", "title", "author", "created_at", "-id", "-title", "-author", "-created_at"}

	if data.ValidateFilters(v, input.Filter); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	books, err := app.models.Books.GetAll(input.Title, input.Author, input.Filter)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
