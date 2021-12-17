package main

import (
	"fmt"
	"log"
	"os"

	"go.fiber.restful/bookmark"
	"go.fiber.restful/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", status)
	app.Get("/api/bookmark", bookmark.GetBookmarks)
	app.Get("/api/bookmark/:id", bookmark.GetBookmark)
	app.Post("/api/bookmark", bookmark.SaveBookmark)
	app.Delete("/api/bookmark/:id", bookmark.DeleteBookmark)
}

func NewServer() *fiber.App {
	app := fiber.New()
	setupRoutes(app)

	return app
}

func main() {
	err := godotenv.Load()
	dbErr := database.InitDatabase()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Print(os.Getenv("postgres_user"))

	if dbErr != nil {
		panic(dbErr)
	}

	app := NewServer()
	app.Listen(":3000")
}
