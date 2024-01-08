package data

import (
	"database/sql"
	"github.com/xuche123/bookwise/internal/validator"
	"time"
)

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ImageURL    string    `json:"image_url,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"-"`
	Version     int32     `json:"version"`
}

func ValidateBook(v *validator.Validator, book *Book) {
	v.Check(book.Title != "", "title", "must be provided")
	v.Check(len(book.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(book.Author != "", "author", "must be provided")
	v.Check(len(book.Author) <= 500, "author", "must not be more than 500 bytes long")
	v.Check(book.ImageURL != "", "image_url", "must be provided")
	v.Check(len(book.ImageURL) <= 500, "image_url", "must not be more than 500 bytes long")
	v.Check(book.Description != "", "description", "must be provided")
	v.Check(len(book.Description) <= 50000, "description", "must not be more than 50000 bytes long")
}

type BookModel struct {
	DB *sql.DB
}

func (m BookModel) Insert(book *Book) error {
	query := `
		INSERT INTO books (title, author, image_url, description) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version`

	args := []any{book.Title, book.Author, book.ImageURL, book.Description}

	return m.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
}

func (m BookModel) Get(id int64) (*Book, error) {
	return nil, nil
}

func (m BookModel) Update(book *Book) error {
	return nil
}

func (m BookModel) Delete(id int64) error {
	return nil
}
