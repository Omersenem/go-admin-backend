package controllers

import (
	"encoding/csv"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/your/repo/database"
	"github.com/your/repo/models"
	"os"
	"strconv"
)

func AllOrders(c *fiber.Ctx) error {
	var filter Filter
	c.BodyParser(&filter)
	pages := strconv.Itoa(filter.Page)
	limit := strconv.Itoa(filter.Limit)

	page, _ := strconv.Atoi(c.Query("page", pages, "limit", limit))
	limitInt, _ := strconv.Atoi(c.Query("limit", limit))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page, limitInt))
}

func Export(c *fiber.Ctx) error {
	filePath := "./csv/orders.csv"

	if err := CreateFile(filePath); err != nil {
		return err
	}

	return c.Download(filePath)
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	for _, order := range orders {
		data := []string{
			strconv.Itoa(int(order.Id)),
			order.FirstName + " " + order.LastName,
			order.Email,
			"",
			"",
			"",
		}

		if err := writer.Write(data); err != nil {
			return err
		}

		for _, orderItem := range order.OrderItems {
			data := []string{
				"",
				"",
				"",
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(int(orderItem.Quantity)),
			}

			if err := writer.Write(data); err != nil {
				return err
			}
		}
	}

	return nil
}

type Sales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

func Chart(c *fiber.Ctx) error {
	var sales []Sales
	fmt.Println(c.Method())

	database.DB.Raw(`
		SELECT DATE_FORMAT(o.created_at, '%Y-%m-%d') as date, SUM(oi.price * oi.quantity) as sum
		FROM orders o
		JOIN order_items oi on o.id = oi.order_id
		GROUP BY date
	`).Scan(&sales)

	return c.JSON(sales)
}
