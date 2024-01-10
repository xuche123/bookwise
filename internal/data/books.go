package data

import (
	"database/sql"
	"errors"
	"fmt"
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
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, title, author, image_url, description, created_at, version
		FROM books
		WHERE id = $1`

	var book Book

	err := m.DB.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ImageURL,
		&book.Description,
		&book.CreatedAt,
		&book.Version,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		} else {
			return nil, err
		}
	}

	return &book, nil
}

func (m BookModel) Update(book *Book) error {
	query := `
		UPDATE books
		SET title = $1, author = $2, image_url = $3, description = $4, version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING version`

	args := []any{book.Title, book.Author, book.ImageURL, book.Description, book.ID, book.Version}

	err := m.DB.QueryRow(query, args...).Scan(&book.Version)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrEditConflict
		} else {
			return err
		}
	}

	return nil
}

func (m BookModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM books
		WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m BookModel) GetAll(title string, author string, filter Filter) ([]*Book, error) {
	query := fmt.Sprintf(`
		SELECT id, title, author, image_url, description, created_at, version
		FROM books
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (to_tsvector('simple', author) @@ plainto_tsquery('simple', $2) OR $2 = '')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filter.sortColumn(), filter.sortDirection())

	rows, err := m.DB.Query(query, title, author, filter.limit(), filter.offset())

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	books := []*Book{}

	for rows.Next() {
		var book Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ImageURL,
			&book.Description,
			&book.CreatedAt,
			&book.Version,
		)

		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
