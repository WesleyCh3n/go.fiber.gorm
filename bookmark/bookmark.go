package bookmark

import "github.com/gofiber/fiber/v2"

func GetAllBookmarks(c *fiber.Ctx) error {
  return c.SendString("All Bookmarks")
}

func SaveBookmark(c *fiber.Ctx) error {
  return c.SendString("Save Bookmarks!")
}
