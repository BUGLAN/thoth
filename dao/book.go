package dao

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/BUGLAN/thoth/model"
)

type BookDao interface {
	FirstOrCreate(book *model.Book) (record *model.Book, err error)
}

type Book struct {
	db *mgo.Database
}

func NewBook(db *mgo.Database) BookDao {
	return &Book{
		db: db,
	}
}

func (dao *Book) FirstOrCreate(book *model.Book) (record *model.Book, err error) {
	record = new(model.Book)
	query := bson.M{"origin_url": book.Origin}
	err = dao.db.C(model.BookC).Find(query).One(record)
	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}

	// first
	if err == mgo.ErrNotFound {
		book.CreateTime, book.UpdateTime = time.Now(), time.Now()
		if err = dao.db.C(model.BookC).Insert(record); err != nil {
			return nil, err
		}
		return book, err
	}

	// update
	book.UpdateTime = time.Now()
	if err = dao.db.C(model.BookC).Update(query, book); err != nil {
		return nil, err
	}
	return book, nil
}

// Chapters save chapters
func (dao *Book) Chapters() {

}
