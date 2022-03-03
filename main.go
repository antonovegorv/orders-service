package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create an instance of the fiber application.
	app := fiber.New()

	// Setup basic handler for a homepage route.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🦆 Quack-quack...")
	})

	// Start the HTTP server to listen for incoming requests.
	log.Fatal(app.Listen(":3000"))
}
