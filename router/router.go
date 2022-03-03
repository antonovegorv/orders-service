package router

import (
	ordersRoutes "github.com/antonovegorv/orders-service/internal/routes/orders"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setups for the entire application.
func SetupRoutes(app *fiber.App) {
	// Create a group with "/api" in route.
	api := app.Group("/api", logger.New())

	// Setup "orders" routes within "/api" route.
	ordersRoutes.SetupOrderRoutes(api)
}
