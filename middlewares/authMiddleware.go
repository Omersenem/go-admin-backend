package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your/repo/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Kimliğiniz doğrulanmadı!",
		})
	}

	return c.Next()
}
