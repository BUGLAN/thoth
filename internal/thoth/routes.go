package thoth

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber"

	"github.com/BUGLAN/thoth/dao"
	"github.com/BUGLAN/thoth/model"
)

type Controller struct {
	book dao.BookDao
}

func NewThothController(book dao.BookDao) *Controller {
	return &Controller{
		book: book,
	}
}

func Route(app *fiber.App, ctrl *Controller) {
	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Status(http.StatusOK).SendString("hello world")
	})
	app.Get("/1", ctrl.Test)
}

func (ctrl *Controller) Test(ctx *fiber.Ctx) {
	book, err := ctrl.book.FirstOrCreate(&model.Book{
		OriginURL:  "test1",
		Book:       "book",
		Name:       "ame",
		Chapters:   []string{"1,", "2"},
		Origin:     "",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		ctx.Status(http.StatusOK).SendString(err.Error())
	}

	ctx.Status(http.StatusOK).Send(book)
}
