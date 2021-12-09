package bookmark

import (
	"github.com/gofiber/fiber/v2"
	"go.fiber.restful/database"
)

func GetAllBookmarks(c *fiber.Ctx) error {
	result, err := database.GetAllBookmarks()
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

func SaveBookmark(c *fiber.Ctx) error {
	newbkmk := new(database.Bookmark)

	err := c.BodyParser(newbkmk)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := database.CreateBookmark(newbkmk)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
