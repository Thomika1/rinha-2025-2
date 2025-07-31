package main

import (
	"log"

	"github.com/Thomika1/rinha-2025-2.git/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	router.InitRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
