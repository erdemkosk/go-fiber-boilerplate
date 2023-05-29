package api

import (
	"go-fiber-boilerplate/api/example"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api")
	example.Routes(v1)
}
