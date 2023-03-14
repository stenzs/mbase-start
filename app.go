package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "mbase/docs"
	"mbase/handlers"
	"mbase/services/dataStorage"
	"mbase/services/database"
)

// mbase
// @title Mbase
// @version 1.0
// @host 0.0.0.0:3000
// @BasePath /

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// Make uploads folder
	dataStorage.MakeUploadFolder()
	database.Engine()
	database.CreateTable()
}

func main() {

	var (
		port = flag.String("port", os.Getenv("PORT"), "Port to listen on")
		prod = flag.Bool("prod", false, "Enable prefork in Production")
	)

	// Parse command-line flags
	flag.Parse()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	allowCredential, _ := strconv.ParseBool(os.Getenv("ALLOW_CREDENTIALS"))
	app.Use(cors.New(cors.Config{
		AllowHeaders:     os.Getenv("ALLOW_HEADERS"),
		AllowOrigins:     os.Getenv("ALLOW_ORIGINS"),
		AllowCredentials: allowCredential,
		AllowMethods:     os.Getenv("GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"),
	}))
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Post("/task", handlers.CreateTask)
	v1.Put("/task", handlers.UpdateTaskStatus)

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
