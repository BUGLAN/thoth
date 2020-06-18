package dao

import (
	"github.com/BUGLAN/thoth/model"
)

type BookDao interface {
}

type Book struct {
}

func NewBook() BookDao {
	return &Book{}
}

func (dao *Book) FirstOrCreate(book *model.Book) (id int64, err error) {
}
