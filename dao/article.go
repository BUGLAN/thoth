package dao

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"

	"github.com/BUGLAN/thoth/model"
)

type ArticleDao interface {
	Create(ctx context.Context, article *model.Article) (err error)
	FirstByID(ctx context.Context, id string) (article *model.Article, err error)
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

func (dao *Article) FirstByID(ctx context.Context, id string) (article *model.Article, err error) {
	article = new(model.Article)
	result, err := dao.es.Get().Index("article").Id(id).Do(ctx)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(result.Source, article); err != nil {
		return nil, err
	}
	return article, err
}
