package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"shoe-store-server/api"
	"shoe-store-server/db"
	"shoe-store-server/helpers"
	"shoe-store-server/helpers/pusher"
	"shoe-store-server/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDatabase()
	//initializers.SeedDB() //have it only run once
	pusher.SetupPusher()
}

func main() {
	// Setup app
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{}))

	// Routes
	api.SetupRoutes(app)


	// Start app
	helpers.HandleError(
		"Error starting application... %s",
		app.Listen(":"+os.Getenv("PORT")),
		true,
	)
}
