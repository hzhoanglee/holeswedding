package router

import (
	"github.com/gofiber/fiber/v2"
	"hz-evite/app/controllers"
)

type HttpRouter struct{}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("")
	group.Get("/", controllers.RenderHello)
	group.Get("/:guestId?", controllers.RenderHello)
	group.Get("/guest/new", controllers.NewUserView)
	group.Post("/new-user", controllers.SaveNewUser)
	group.Post("/words", controllers.SaveWords)
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
