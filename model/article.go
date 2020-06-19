package model

type Article struct {
	Content    string `json:"content"`
	Title      string `json:"title"`
	CreateTime string `json:"create_time"`
}
