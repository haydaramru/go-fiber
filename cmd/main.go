package main

import (
	"github.com/gofiber/fiber/v2"
)

func check(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New()

	app.Get("/", check)

	app.Listen(":3000")
}
