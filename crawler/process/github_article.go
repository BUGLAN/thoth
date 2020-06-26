package process

import (
	"context"
	"fmt"
	"time"

	"github.com/BUGLAN/thoth/crawler"
	"github.com/BUGLAN/thoth/model"
)

// GithubReadMeProcess
func (process *KanShu) GithubReadMeProcess(ctx *crawler.Context) error {
	fmt.Println(ctx.Html)
	title := "java Guide"
	err := process.article.Create(context.Background(), &model.Article{
		Content:    ctx.Html,
		Title:      title,
		CreateTime: time.Now().String(),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
