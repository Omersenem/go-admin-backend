package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your/repo/database"
	"github.com/your/repo/models"
	"strconv"
)

func AllProducts(c *fiber.Ctx) error {
	var filter Filter
	c.BodyParser(&filter)
	pages := strconv.Itoa(filter.Page)
	limit := strconv.Itoa(filter.Limit)

	page, _ := strconv.Atoi(c.Query("page", pages, "limit", limit))
	limitInt, _ := strconv.Atoi(c.Query("limit", limit))

	return c.JSON(models.Paginate(database.DB, &models.Product{}, page, limitInt))
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return nil
}
