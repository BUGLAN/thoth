package dao

import (
	"github.com/globalsign/mgo"

	"github.com/BUGLAN/thoth/model"
)

type ArticleDao interface {
	FirstOrCreate(book *model.Book) (record *model.Book, err error)
}

type Article struct {
	db *mgo.Database
}