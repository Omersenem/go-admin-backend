package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/your/repo/database"
	"github.com/your/repo/middlewares"
	"github.com/your/repo/models"
	"strconv"
)

type Filter struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func AllUsers(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var filter Filter
	c.BodyParser(&filter)
	pages := strconv.Itoa(filter.Page)
	limit := strconv.Itoa(filter.Limit)

	page, _ := strconv.Atoi(c.Query("page", pages, "limit", limit))
	limitInt, _ := strconv.Atoi(c.Query("limit", limit))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page, limitInt))
}

func CreateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	fmt.Println("İddd", id)

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
