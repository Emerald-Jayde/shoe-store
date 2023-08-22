package api

import (
	"github.com/gofiber/fiber/v2"
	"shoe-store-server/api/v1"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/stores", v1.GetStores)
	api.Get("/stores/:id", v1.GetStore)

	api.Get("/shoe_models", v1.GetShoeModels)
	api.Get("/shoe_models/:id", v1.GetShoeModel)
	api.Post("/shoe_models", v1.CreateShoeModel)

	api.Get("/inventory", v1.GetAllInventory)
	api.Get("/inventory/store/:id", v1.GetInventoryForStore)
	api.Get("/inventory/suggestions", v1.InventoryMoveSuggestions)

	api.Get("/sales", v1.GetSales)
	api.Get("/sales/store/:id", v1.GetSalesForStore)
	api.Get("/sales/limit/:id", v1.GetLastXSales)

	app.Use(NotFoundRoute)
}

// NotFoundRoute func to describe 404 Error route.
func NotFoundRoute(c *fiber.Ctx) error {
	// Return HTTP 404 status and JSON response.
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": true,
		"msg":   "sorry, endpoint not found",
	})
}
