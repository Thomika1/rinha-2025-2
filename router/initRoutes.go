package router

import (
	"github.com/Thomika1/rinha-2025-2.git/handler"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("/Payments", handler.Payments)
}
