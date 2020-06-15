/*
web应用的入口
*/
package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber"
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
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	App(func() {
		log.Fatal(app.Listen(8080))
	})
}
