package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/your/repo/database"
	"github.com/your/repo/routes"
)

func main() {

	// Connect Database
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8080",
	}))

	//Routes
	routes.Setup(app)

	// Start the server on port 8080
	app.Listen(":8080")

}
