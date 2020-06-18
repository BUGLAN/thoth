/*
web应用的入口
*/
package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber"

	"github.com/BUGLAN/thoth/internal/thoth"
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
	app := fiber.New()

	App(func() {
		thothController, err := wire.InitThothController()
		if err != nil {
			panic(err)
		}

		thoth.Route(app, thothController)
		log.Fatal(app.Listen(8080))
	})
}
