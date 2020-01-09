package model

import "time"

type (
	Post struct {
		FileName   string     `json:"objectID"`
		Header     PostHeader `json:"header"`
		RowContent string     `json:"content"`
		Content    string     `json:"-"`
	}

	PostHeader struct {
		Layout     string    `json:"layout"`
		Title      string    `json:"title"`
		Date       time.Time `json:"date"`
		Comments   string    `json:"comments"`
		Categories []string  `json:"categories"`
		Tags       []string  `json:"tags"`
	}
)
