// +build wireinject

package wire

import (
	"github.com/globalsign/mgo"
	"github.com/google/wire"
	"github.com/olivere/elastic/v7"

	"github.com/BUGLAN/thoth/crawler/process"
	"github.com/BUGLAN/thoth/dao"
	"github.com/BUGLAN/thoth/internal/pkg/es"
	"github.com/BUGLAN/thoth/internal/pkg/mongo"
	"github.com/BUGLAN/thoth/internal/thoth/service"
)

func InitMongo() *mgo.Database {
	return mongo.Init(&mongo.Config{
		Username:   "root",
		Password:   "root",
		ServerList: "localhost:27017",
		Database:   "admin",
		Option:     "",
		PoolLimit:  20,
	})

}

func InitEs() *elastic.Client {
	return es.Init(&es.Config{ServerURLs: []string{"http://127.0.0.1:9200"}})
}

func InitBookDao() dao.BookDao {
	panic(wire.Build(InitMongo, dao.NewBook))
}

func InitArticleDao() dao.ArticleDao {
	panic(wire.Build(InitEs, dao.NewArticle))
}

func InitArticleService() (article service.ArticleService) {
	panic(wire.Build(InitEs, dao.NewArticle, service.NewArticleService))

}

// func InitThothController() (*thoth.Controller, error) {
// 	panic(wire.Build(InitBookDao, thoth.NewThothController))
// }
//
func InitKanShuProcess() *process.KanShu {
	panic(wire.Build(InitBookDao, InitArticleDao, process.NewKanShu))

}
