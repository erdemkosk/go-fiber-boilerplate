package server

import (
	"fmt"
	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/swagger"

	_ "go-fiber-boilerplate/docs"
)

var cfg config.Config = config.LoadConfig()

func setupMiddlewares(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())
	app.Use(requestid.New())

	if cfg.ENABLE_LIMITER == true {
		app.Use(limiter.New(limiter.Config{
			Max:               20,
			Expiration:        30 * time.Second,
			LimiterMiddleware: limiter.SlidingWindow{},
		}))
	}

	if cfg.ENABLE_LOGGER == true {
		app.Use(logger.New(logger.Config{
			// For more options, see the Config section
			Format: "[${ip}]:${port} ${locals:requestid} ${status} - ${method} ${path}​\n",
		}))
	}
}

func Create() *fiber.App {
	//database.SetupDatabase()

	const idleTimeout = 5 * time.Second

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*pkg.Error); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*fiber.Error); ok {
				return ctx.Status(e.Code).JSON(pkg.Error{Status: e.Code, Code: "internal-server", Message: e.Message})
			} else {
				return ctx.Status(500).JSON(pkg.Error{Status: 500, Code: "internal-server", Message: err.Error()})
			}
		},
		IdleTimeout: idleTimeout,
	})

	setupMiddlewares(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	if cfg.ENABLE_MONITOR == true {
		app.Get("/metrics", monitor.New())
	}

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	return app
}

func Listen(app *fiber.App) error {

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return app.Listen(fmt.Sprintf("%s:%s", cfg.SERVER_HOST, cfg.SERVER_PORT))
}
