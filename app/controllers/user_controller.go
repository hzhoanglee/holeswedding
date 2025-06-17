package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hz-evite/app/models"
	"hz-evite/pkg/database"
)

func NewUserView(c *fiber.Ctx) error {
	c.Set("WWW-Authenticate", `Basic realm="Super User"`)
	if c.Get("Authorization") != "Basic aHpob2FuZ2xlZTpIb2FuZ0xlMjgwNg==" { // user:password
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	return c.Render("super-new", fiber.Map{})
}

func SaveNewUser(c *fiber.Ctx) error {
	name := c.FormValue("user")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Name is required")
	}
	phone := c.FormValue("phone")
	if phone == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Phone is required")
	}
	uID, _ := uuid.NewV7()
	userModel := models.User{}
	userModel.ID = uID
	userModel.Name = name
	userModel.Phone = phone
	database.DB.Create(&userModel)
	if userModel.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create user")
	}
	return c.JSON(&userModel)

}

func SaveWords(c *fiber.Ctx) error {
	from := c.FormValue("Name")
	message := c.FormValue("Message")
	if from == "" || message == "" {
		return c.Status(fiber.StatusBadRequest).SendString("From and Message are required")
	}
	wordObject := models.Words{
		From:    from,
		Message: message,
	}
	database.DB.Create(&wordObject)
	return c.JSON(&wordObject)
}
