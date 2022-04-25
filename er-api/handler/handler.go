package handler

import (
	"er-api/db"
	"er-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetExchanges(c *fiber.Ctx) error {
	db := db.Init()
	var exchanges []models.Exchanges
	db.Find(&exchanges)
	if len(exchanges) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No exchanges present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Exchanges Found", "data": exchanges})
}

func GetExchange(c *fiber.Ctx) error {
	db := db.Init()
	var exchange models.Exchanges
	name := c.Params("name")
	db.Last(&exchange)
	if exchange.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No exchange present", "data": nil})
	}

	switch name {
	case "usd", "USD":
		exchange.USD = 1.0
		return c.JSON(fiber.Map{"status": "success", "message": "Exchange Found", "data": exchange})
	case "eur", "EUR":
		exchange.USD = (exchange.TRY / exchange.EUR) - exchange.TRY
		exchange.TRY = exchange.TRY / exchange.EUR
		exchange.EUR = 1.0
		return c.JSON(fiber.Map{"status": "success", "message": "Exchange Found", "data": exchange})
	case "try", "TRY":
		exchange.USD = exchange.TRY
		exchange.EUR = exchange.TRY / exchange.EUR
		exchange.TRY = 1
		return c.JSON(fiber.Map{"status": "success", "message": "Exchange Found", "data": exchange})
	default:
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No exchange present", "data": nil})
	}
}
