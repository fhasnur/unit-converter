package converter

import (
	"strconv"

	"github.com/fhasnur/unit-converter/internal/data"
	"github.com/gofiber/fiber/v2"
)

var weightFactors = map[string]float64{
	"milligram": 0.000001,
	"gram":      0.001,
	"kilogram":  1.0,
	"ounce":     0.0283495,
	"pound":     0.453592,
}

func ConvertWeight(ctx *fiber.Ctx) error {
	valueStr := ctx.FormValue("value")
	fromUnit := ctx.FormValue("from")
	toUnit := ctx.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid value")
	}

	fromFactor, fromExists := weightFactors[fromUnit]
	toFactor, toExists := weightFactors[toUnit]

	if !fromExists || !toExists {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid units")
	}

	convertedValue := (value * fromFactor) / toFactor

	weightResult := fiber.Map{
		"InputValue":     value,
		"FromUnit":       fromUnit,
		"ToUnit":         toUnit,
		"ConvertedValue": convertedValue,
	}

	dataUnits := data.GetDataUnits(nil, weightResult, nil)

	return ctx.Render("index", fiber.Map{
		"Converters":        dataUnits,
		"LengthResult":      nil,
		"WeightResult":      weightResult,
		"TemperatureResult": nil,
	})
}
