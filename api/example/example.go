package example

import (
	"github.com/gofiber/fiber/v2"
)

type Example struct {
	Name string `json:"name"`
}

func GetExamples(c *fiber.Ctx) error {
	var examples = []Example{
		{Name: "Erdem"},
		{Name: "Köşk"},
	}

	return c.JSON(examples)
}
