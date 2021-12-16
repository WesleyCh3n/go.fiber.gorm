package bookmark

import (
	"github.com/gofiber/fiber/v2"
	"go.fiber.restful/database"
)

func GetBookmarks(c *fiber.Ctx) error {
	result, err := database.GetBookmarks()
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

func GetBookmark(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := database.GetBookmark(id)
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
		"ID":      newbkmk.ID,
	})
	return nil
}

func DeleteBookmark(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DeleteBookmark(id)

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
		"message": "bookmark deleted",
		"data":    nil,
	})

	return nil
}
