package main

import (
	"fmt"
	"log"

	"example.com/learn-golang/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func check(c *fiber.Ctx) error {
	return c.SendString("And so here the adventure begins")
}

func main() {
	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Hello from middleware")
		return c.Next()
	})

	app.Get("/", check)

	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProducts)

	log.Fatal(app.Listen(":3000"))
}
