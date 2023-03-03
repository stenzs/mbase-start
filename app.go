package main

import (
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "mbase/docs"
	"mbase/handlers"
)

// mbase
// @version 1.0
// @host localhost:3000
// @BasePath /

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	portNumber, _ := os.LookupEnv("PORT")
	var (
		port = flag.String("port", portNumber, "Port to listen on")
		prod = flag.Bool("prod", false, "Enable prefork in Production")
	)

	// Parse command-line flags
	flag.Parse()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Post("/update", handlers.UpdateData)

	// Setup static files
	app.Static("/", "./static/public")

	// Swagger
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}

// swaggerDock godoc
// @Router /api/v1/users [get]
// @Router /api/v1/users [post]
//func swaggerDockz() {}
