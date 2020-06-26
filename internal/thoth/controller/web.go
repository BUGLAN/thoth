package controller

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"

	"github.com/BUGLAN/thoth/internal/thoth/service"
)

type Controller struct {
	articleService service.ArticleService
}

// NewThothController thoth控制器
func NewThothController(view *html.Engine, app *fiber.App, articleService service.ArticleService) {
	ctrl := &Controller{articleService: articleService}

	app.Get("/", func(ctx *fiber.Ctx) { ctx.Status(http.StatusOK).SendString("hello world") })
	app.Get("/index", ctrl.Index)
	view.AddFunc("unescaped", func(html string) interface{} {
		return template.HTML(html)
	})
}

// Index 首页
func (ctrl *Controller) Index(ctx *fiber.Ctx) {
	var err error
	id := ctx.Query("id")
	article, err := ctrl.articleService.FirstArticleByID(context.Background(), id)
	if err != nil {
		panic(err)
	}

	err = ctx.Render("index", fiber.Map{
		"title":        article.Title,
		"content":      article.Content,
		"updated_time": article.CreateTime,
	}, "index")
	if err != nil {
		fmt.Println(err.Error())
	}
}
