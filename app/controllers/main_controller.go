package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hz-evite/app/models"
	"hz-evite/pkg/database"
)

func RenderHello(c *fiber.Ctx) error {
	guestId := c.Params("guestId")
	if guestId == "" {
		guestId = "00000000-0000-0000-0000-000000000000"
	}
	user := models.User{}
	db := database.DB
	if err := db.First(&user, "id = ?", guestId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.Render("index", fiber.Map{
		"guestName": user.Name,
	})
}
