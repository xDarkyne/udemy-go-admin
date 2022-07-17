package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xdarkyne/udemy/config"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/routes"
)

func main() {
	config.LoadConfig()
	db.Connect()
	db.Migrate()

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")
}
