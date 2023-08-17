package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"shoe-store-server/helpers"

	"os"
	"shoe-store-server/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.WebsocketClient()
}

func main() {
	// Setup app
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{}))

	// Routes
	Routes(app)

	// Start app
	helpers.HandleError(
		"Error starting application... %s",
		app.Listen(":"+os.Getenv("PORT")),
		true,
	)
}
