package controller

import (
	"onestep/database"
	"onestep/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Preload("Role").Find(&users)
	return c.Status(fiber.StatusOK).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("asd"), 10)

	user.Password = password

	database.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var user models.User

	database.DB.Preload("Role").First(&user, id)

	return c.Status(fiber.StatusOK).JSON(user)

}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Delete(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Deleted",
	})
}
