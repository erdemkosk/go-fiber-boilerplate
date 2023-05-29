package example

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Get("/examples", GetExamples)
	route.Post("/examples", CreateExample)
}
