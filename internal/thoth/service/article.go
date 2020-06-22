package service

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"

	"github.com/BUGLAN/thoth/model"
)

const article = "article"

type IArticleService interface {
}

type ArticleService struct {
	es *elastic.Client
}

func NewArticleService(es *elastic.Client) *ArticleService {
	return &ArticleService{
		es: es,
	}
}

// CreateArticle 创建或修改es信息(update==delete)
func (svc ArticleService) CreateOrUpdateArticle(ctx context.Context, article *model.Article) (err error) {
	_, err = svc.es.Index().Index("article").BodyJson(article).Do(ctx)
	return err
}

// SearchArticle 搜索es信息(更具标题和文章内容)
func (svc ArticleService) SearchArticle(ctx context.Context, s string) (articles []*model.Article, err error) {
	query := elastic.NewBoolQuery()
	query = query.Should(elastic.NewMatchQuery("title", s))
	query = query.Should(elastic.NewMatchQuery("content", s))
	result, err := svc.es.Search().Index(article).Query(query).Do(ctx)

	var article model.Article

	if result == nil || len(result.Hits.Hits) == 0 {
		return
	}

	for _, hit := range result.Hits.Hits {
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			return
		}
		articles = append(articles, &article)
	}

	return articles, err
}
