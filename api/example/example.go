package example

import (
	"go-fiber-boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
)

type Example struct {
	Name string `json:"name"`
}

// Login is a function to get a book by ID
// @Summary Get a all examples
// @Description Get all examples
// @Tags examples
// @Accept json
// @Produce json
// @Success 200 {array} Example[]
// @Router /api/examples [get]
func GetExamples(c *fiber.Ctx) error {
	var examples = []Example{
		{Name: "Erdem"},
		{Name: "Köşk"},
	}

	return c.JSON(examples)
}

// Login is a function to get a book by ID
// @Summary Create a new example
// @Description Create a new example
// @Tags examples
// @Accept json
// @Produce json
// @Param cridentinials body Example true "Example"
// @Success 200 {object} Example
// @Router /api/examples [post]
func CreateExample(c *fiber.Ctx) error {
	example := new(Example)
	if err := c.BodyParser(example); err != nil {
		return pkg.BadRequest("Invalid params")
	}

	return c.JSON(example)
}
