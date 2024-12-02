package main

import (
	"github.com/fhasnur/unit-converter/internal/converter"
	"github.com/fhasnur/unit-converter/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	engine := mustache.New("./templates", ".mustache")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", handler.RootHandler)
	app.Post("/length", converter.ConvertLength)
	app.Post("/weight", converter.ConvertWeight)
	app.Post("/temperature", converter.ConvertTemperature)

	app.Static("/static", "./static")

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
