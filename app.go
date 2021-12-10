package main

import (
	"go.fiber.restful/bookmark"
	"go.fiber.restful/database"

	"github.com/gofiber/fiber/v2"
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
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	app := NewServer()
	app.Listen(":3000")
}
