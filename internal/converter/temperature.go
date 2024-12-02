package converter

import (
	"strconv"

	"github.com/fhasnur/unit-converter/internal/data"
	"github.com/gofiber/fiber/v2"
)

func ConvertTemperature(ctx *fiber.Ctx) error {
	valueStr := ctx.FormValue("value")
	fromUnit := ctx.FormValue("from")
	toUnit := ctx.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid value")
	}

	var convertedValue float64
	switch {
	case fromUnit == "celsius" && toUnit == "fahrenheit":
		convertedValue = (value * 9 / 5) + 32
	case fromUnit == "celsius" && toUnit == "kelvin":
		convertedValue = value + 273.15
	case fromUnit == "fahrenheit" && toUnit == "celsius":
		convertedValue = (value - 32) * 5 / 9
	case fromUnit == "fahrenheit" && toUnit == "kelvin":
		convertedValue = (value-32)*5/9 + 273.15
	case fromUnit == "kelvin" && toUnit == "celsius":
		convertedValue = value - 273.15
	case fromUnit == "kelvin" && toUnit == "fahrenheit":
		convertedValue = (value-273.15)*9/5 + 32
	case fromUnit == toUnit:
		convertedValue = value
	default:
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid units")
	}

	temperatureResult := fiber.Map{
		"InputValue":     value,
		"FromUnit":       fromUnit,
		"ToUnit":         toUnit,
		"ConvertedValue": convertedValue,
	}

	dataUnits := data.GetDataUnits(nil, nil, temperatureResult)

	return ctx.Render("index", fiber.Map{
		"Converters":        dataUnits,
		"LengthResult":      nil,
		"WeightResult":      nil,
		"TemperatureResult": temperatureResult,
	})
}
