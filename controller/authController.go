package controller

import (
	"onestep/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		Firstname: "John",
		Lastname:  "Wick",
		Email:     "john@wick.com",
		Password:  "xxxx",
	}

	return c.JSON(user)
}
