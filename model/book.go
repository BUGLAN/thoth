package model

import (
	"time"
)

const BookC = "book"

type Book struct {
	OriginURL  string    `json:"origin_url" bson:"origin_url"`
	Book       string    `json:"book" bson:"book"`
	Name       string    `json:"name" bson:"name"`
	Chapters   []string  `json:"chapters" bson:"chapters"`
	Origin     string    `json:"origin" bson:"origin"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}
