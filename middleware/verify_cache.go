package middleware

import (
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/antonovegorv/orders-service/cache"
	"github.com/gofiber/fiber/v2"
)

// VerifyCache checks if there is an order with the given id in the cache. If it
// is then it returns cached value. Otherwise, it passes execution to the next
// middleware.
func VerifyCache(c *fiber.Ctx) error {
	// Grab an order id from the route.
	id := c.Params("orderId")

	// Check the value in the cache.
	val, err := cache.Orders.Get(id)
	if err != ttlcache.ErrNotFound {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "An order was found in cache!",
			"data":    val,
		})
	}

	// Go to next middleware.
	return c.Next()
}
