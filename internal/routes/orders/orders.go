package ordersRoutes

import (
	ordersHandler "github.com/antonovegorv/orders-service/internal/handlers/orders"
	"github.com/antonovegorv/orders-service/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupOrderRoutes setups all the endpoints for the "/orders" route.
func SetupOrderRoutes(router fiber.Router) {
	// Create "/orders" router group.
	orders := router.Group("/orders")

	// Read an order by id.
	orders.Get("/:orderId", middleware.VerifyCache, ordersHandler.GetOrder)
}
