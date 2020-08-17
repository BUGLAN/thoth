/*
web应用的入口
*/
package main

import (
	"github.com/gofiber/fiber/middleware"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"

	"github.com/BUGLAN/thoth/internal/thoth/controller"
	"github.com/BUGLAN/thoth/wire"
)

func App(f func()) {
	quit := make(chan os.Signal, 0)
	signal.Notify(quit, os.Interrupt)

	go f()

	<-quit
	log.Println("\nSTOP!!!")
}

func main() {
	views := html.New("./web/template", ".html")
	views.Reload(true)
	// views.Delims("{{", "}}")
	app := fiber.New(&fiber.Settings{Views: views})
	app.Use(middleware.Recover())

	App(func() {
		service := wire.InitArticleService()
		controller.NewThothController(views, app, service)

		log.Fatal(app.Listen(8080))
	})
}
