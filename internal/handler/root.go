package handler

import (
	"github.com/fhasnur/unit-converter/internal/data"
	"github.com/gofiber/fiber/v2"
)

func RootHandler(ctx *fiber.Ctx) error {
	lengthResult := fiber.Map{
		"InputValue":     nil,
		"FromUnit":       nil,
		"ToUnit":         nil,
		"ConvertedValue": nil,
	}
	weightResult := fiber.Map{
		"InputValue":     nil,
		"FromUnit":       nil,
		"ToUnit":         nil,
		"ConvertedValue": nil,
	}
	temperatureResult := fiber.Map{
		"InputValue":     nil,
		"FromUnit":       nil,
		"ToUnit":         nil,
		"ConvertedValue": nil,
	}

	dataUnits := data.GetDataUnits(lengthResult, weightResult, temperatureResult)

	return ctx.Render("index", fiber.Map{
		"Converters":        dataUnits,
		"LengthResult":      lengthResult,
		"WeightResult":      weightResult,
		"TemperatureResult": temperatureResult,
	})
}
