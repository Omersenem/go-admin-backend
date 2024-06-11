package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your/repo/database"
	"github.com/your/repo/routes"
)

func main() {

	// Connect Database
	database.Connect()

	app := fiber.New()

	//Routes
	routes.Setup(app)

	// Start the server on port 8080
	app.Listen(":8080")

}
