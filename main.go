package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	engine := mustache.New("./templates", ".mustache")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Result": "",
		})
	})

	app.Static("/static", "./static")

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
