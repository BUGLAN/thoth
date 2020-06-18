package dao

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/BUGLAN/thoth/model"
)

type BookDao interface {
}

type Book struct {
	session *mgo.Session
	mgo     *mgo.Database
}

func NewBook() BookDao {
	return &Book{}
}

func (dao *Book) FirstOrCreate(book *model.Book) (record *model.Book, err error) {
	record = new(model.Book)
	query := bson.M{"origin_url": book.Origin}
	err = dao.mgo.C(model.BookC).Find(query).One(record)
	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}

	// first
	if err == mgo.ErrNotFound {
		book.CreateTime, book.UpdateTime = time.Now(), time.Now()
		if err = dao.mgo.C(model.BookC).Insert(record); err != nil {
			return nil, err
		}
		return book, err
	}

	// update
	book.UpdateTime = time.Now()
	if _, err = dao.mgo.C(model.BookC).Upsert(query, book); err != nil {
		return nil, err
	}
	return book, nil
}
