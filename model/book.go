package model

import (
	"time"
)

const BookC = "book"

type Book struct {
	OriginURL  string    `json:"origin_url"`
	Book       string    `json:"book"`
	Name       string    `json:"name"`
	Chapters   []string  `json:"chapters"`
	Origin     string    `json:"origin"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
