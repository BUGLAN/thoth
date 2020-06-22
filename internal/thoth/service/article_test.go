package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/BUGLAN/thoth/model"
	"github.com/BUGLAN/thoth/wire"
)

var (
	svc *ArticleService
	ctx = context.Background()
	err error
)

func init() {
	es, err := wire.InitEs()
	if err != nil {
		panic(err)
	}
	svc = NewArticleService(es)
}

func TestArticleService_CreateOrUpdateArticle(t *testing.T) {
	if svc.es == nil {
		t.Fatal("es is nil")
	}

	err = svc.CreateOrUpdateArticle(ctx,
		&model.Article{
			Title:   "大师",
			Content: "这是一个content",
		})

	if err != nil {
		t.Error(err)
	}
}

func TestArticleService_SearchArticle(t *testing.T) {
	articles, err := svc.SearchArticle(ctx, "content2")
	assert.Nil(t, err)
	assert.NotZero(t, len(articles))
	for _, article := range articles {
		fmt.Println(article.Content)
		fmt.Println(article.Title)
	}
}
