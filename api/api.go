package api

import (
	"go-fiber-boilerplate/api/example"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	example.Routes(v1)
}
