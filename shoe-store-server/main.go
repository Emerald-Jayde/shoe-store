package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"shoe-store-server/api"
	"shoe-store-server/db"
	"shoe-store-server/initializers"
	"shoe-store-server/lib/pusher"
	"shoe-store-server/lib/websocket"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDatabase()
	//initializers.SeedDB() //have it only run once
	pusher.SetupPusher()
}

func main() {
	addr := flag.String("addr", ":"+os.Getenv("PORT"), "http service address")

	// Setup app
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{}))

	// Routes
	api.SetupRoutes(app)

	// Start websocket client
	_ = ws.StartWebsocket(os.Getenv("WEBSOCKET_HOST"), "")

	// Start app
	if err := app.Listen(*addr); err != nil {
		log.Fatalf("Error starting application... %s", err)
	}
}
