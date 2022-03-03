package ordersHandler

import (
	"github.com/gofiber/fiber/v2"
)

// GetOrder is a placeholder. It returns a json with a message, that we were
// unable to find an order in cache.
// P.S. In real application here we should implement some logic to make a
// request to the database and try to search there. Otherwise, we should reply
// with an error message. But for this application we should not do this.
func GetOrder(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  "error",
		"message": "An order was not found in cache :/",
		"data":    nil,
	})
}
