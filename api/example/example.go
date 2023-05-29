package example

import (
	"go-fiber-boilerplate/pkg"

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

func CreateExample(c *fiber.Ctx) error {
	example := new(Example)
	if err := c.BodyParser(example); err != nil {
		return pkg.BadRequest("Invalid params")
	}

	return c.JSON(example)
}
