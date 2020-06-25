package dao

import (
	"context"

	"github.com/olivere/elastic/v7"

	"github.com/BUGLAN/thoth/model"
)

type ArticleDao interface {
	Create(ctx context.Context, article *model.Article) (err error)
}

type Article struct {
	es *elastic.Client
}

func NewArticle(es *elastic.Client) ArticleDao {
	return &Article{
		es: es,
	}
}

func (dao *Article) Create(ctx context.Context, article *model.Article) (err error) {
	_, err = dao.es.Index().Index("article").BodyJson(article).Do(ctx)
	return
}
