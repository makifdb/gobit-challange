package main

import (
	"er-api/db"
	"er-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db.Init()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
