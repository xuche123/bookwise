package data

import "time"

type Book struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date,omitempty"`
	ImageURL      string    `json:"image_url,omitempty"`
	Description   string    `json:"description,omitempty"`
	CreatedAt     time.Time `json:"-"`
	Version       int32     `json:"version"`
}
