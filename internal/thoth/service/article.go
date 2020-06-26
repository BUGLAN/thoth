package service

import (
	"context"

	"github.com/BUGLAN/thoth/dao"
	"github.com/BUGLAN/thoth/model"
)

type ArticleService interface {
	CreateOrUpdateArticle(ctx context.Context, article *model.Article) (err error)
	SearchArticle(ctx context.Context, s string) (articles []*model.Article, err error)
	FirstArticleByID(ctx context.Context, id string) (article *model.Article, err error)
}

type articleService struct {
	esDao dao.ArticleDao
}

func NewArticleService(esDao dao.ArticleDao) ArticleService {
	return &articleService{
		esDao: esDao,
	}
}

// CreateArticle 创建或修改es信息(update==delete)
func (svc *articleService) CreateOrUpdateArticle(ctx context.Context, article *model.Article) (err error) {
	return svc.esDao.Create(ctx, article)
}

// SearchArticle 搜索es信息(更具标题和文章内容)
func (svc *articleService) SearchArticle(ctx context.Context, s string) (articles []*model.Article, err error) {
	// query := elastic.NewBoolQuery()
	// query = query.Should(elastic.NewMatchQuery("title", s))
	// query = query.Should(elastic.NewMatchQuery("content", s))
	// result, err := svc.es.Search().Index(article).Query(query).Do(ctx)
	//
	// var article model.Article
	//
	// if result == nil || len(result.Hits.Hits) == 0 {
	// 	return
	// }
	//
	// for _, hit := range result.Hits.Hits {
	// 	err = json.Unmarshal(hit.Source, &article)
	// 	if err != nil {
	// 		return
	// 	}
	// 	articles = append(articles, &article)
	// }

	return articles, err
}

func (svc *articleService) FirstArticleByID(ctx context.Context, id string) (article *model.Article, err error) {
	return svc.esDao.FirstByID(ctx, id)
}
