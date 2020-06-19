// +build wireinject

package wire

import (
	"github.com/globalsign/mgo"
	"github.com/google/wire"

	"github.com/BUGLAN/thoth/dao"
	"github.com/BUGLAN/thoth/internal/pkg/mongo"
	"github.com/BUGLAN/thoth/internal/thoth"
)

func InitMongoConfig() (*mgo.Database, error) {
	return mongo.Init(&mongo.Config{
		Username:   "root",
		Password:   "root",
		ServerList: "localhost:27017",
		Database:   "admin",
		Option:     "",
		PoolLimit:  20,
	})

}

func InitBookDao() (dao.BookDao, error) {
	panic(wire.Build(InitMongoConfig, dao.NewBook))
}

func InitThothController() (*thoth.Controller, error) {
	panic(wire.Build(InitBookDao, thoth.NewThothController))
}
