package router

import (
	"er-api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	exc := api.Group("/exchange")
	exc.Get("/", handler.GetExchanges)
	exc.Get("/:name", handler.GetExchange)
}
