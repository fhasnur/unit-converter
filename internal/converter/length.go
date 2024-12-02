package converter

import (
	"strconv"

	"github.com/fhasnur/unit-converter/internal/data"
	"github.com/gofiber/fiber/v2"
)

var lengthFactors = map[string]float64{
	"millimeter": 0.001,
	"centimeter": 0.01,
	"meter":      1.0,
	"kilometer":  1000.0,
	"inch":       0.0254,
	"foot":       0.3048,
	"yard":       0.9144,
	"mile":       1609.34,
}

func ConvertLength(ctx *fiber.Ctx) error {
	valueStr := ctx.FormValue("value")
	fromUnit := ctx.FormValue("from")
	toUnit := ctx.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid value")
	}

	fromFactor, fromExists := lengthFactors[fromUnit]
	toFactor, toExists := lengthFactors[toUnit]
	if !fromExists || !toExists {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid units")
	}

	convertedValue := (value * fromFactor) / toFactor

	lengthResult := fiber.Map{
		"InputValue":     value,
		"FromUnit":       fromUnit,
		"ToUnit":         toUnit,
		"ConvertedValue": convertedValue,
	}

	dataUnits := data.GetDataUnits(lengthResult, nil, nil)

	return ctx.Render("index", fiber.Map{
		"Converters":        dataUnits,
		"LengthResult":      lengthResult,
		"WeightResult":      nil,
		"TemperatureResult": nil,
	})
}
