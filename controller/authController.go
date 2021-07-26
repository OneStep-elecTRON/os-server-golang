package controller

import (
	"onestep/database"
	"onestep/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	if data["password"] != data["conf_password"] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "password does not match",
		})

	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	user := models.User{
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	database.DB.Create(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user *models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "successful login",
	})

}
