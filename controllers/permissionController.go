package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your/repo/database"
	"github.com/your/repo/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
