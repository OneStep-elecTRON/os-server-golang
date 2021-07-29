package controller

import (
	"onestep/database"
	"onestep/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)
	return c.Status(fiber.StatusOK).JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Create(&role)

	return c.Status(fiber.StatusCreated).JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var role models.Role

	database.DB.First(&role, id)

	return c.Status(fiber.StatusOK).JSON(role)

}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		ID: uint(id),
	}

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Model(&role).Updates(role)

	return c.Status(fiber.StatusOK).JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		ID: uint(id),
	}

	database.DB.Delete(&role)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "role Deleted",
	})
}
