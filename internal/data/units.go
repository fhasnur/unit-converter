package data

import "github.com/gofiber/fiber/v2"

func GetDataUnits(lengthResult, weightResult, temperatureResult fiber.Map) []fiber.Map {
	return []fiber.Map{
		{
			"Id":      "length",
			"Title":   "Length",
			"Action":  "/length",
			"Options": []string{"millimeter", "centimeter", "meter", "kilometer", "inch", "foot", "yard", "mile"},
			"Result":  lengthResult,
			"active":  true,
		},
		{
			"Id":      "weight",
			"Title":   "Weight",
			"Action":  "/weight",
			"Options": []string{"milligram", "gram", "kilogram", "ounce", "pound"},
			"Result":  weightResult,
			"active":  false,
		},
		{
			"Id":      "temperature",
			"Title":   "Temperature",
			"Action":  "/temperature",
			"Options": []string{"celsius", "fahrenheit", "kelvin"},
			"Result":  temperatureResult,
			"active":  false,
		},
	}
}
