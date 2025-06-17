package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"hz-evite/pkg/database"
	"hz-evite/pkg/env"
	"hz-evite/pkg/router"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	database.SetupDatabase()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		}))
	app.Use(logger.New(
		logger.Config{}))
	app.Get("/dashboard", monitor.New())
	app.Static("/", "./assets")
	router.InstallRouter(app)

	return app
}
