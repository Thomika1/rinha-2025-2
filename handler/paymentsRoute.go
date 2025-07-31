package handler

import (
	"github.com/Thomika1/rinha-2025-2.git/model"
	"github.com/gofiber/fiber/v2"
)

func Payments(ctx *fiber.Ctx) error {

	var payment model.Payments

	err := ctx.BodyParser(&payment)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "could not parse JSON",
		})
	}

	return ctx.Status(fiber.StatusAccepted).SendString("payment queued")
}
