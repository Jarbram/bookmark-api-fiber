package services

import (
	"bookmark-api-fiber/database"
	"bookmark-api-fiber/models"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	database *database.Database
}

func NewService(database *database.Database) *Service {
	return &Service{database}
}

func (s *Service) Status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func (s *Service) GetAllBookmarks(c *fiber.Ctx) error {
	result, err := s.database.GetAllBookmarks()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

func (s *Service) GetBookmarkId(c *fiber.Ctx) error {

	id := c.Params("id")

	result, err := s.database.GetBookmarkId(id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

func (s *Service) SaveBookmark(c *fiber.Ctx) error {
	newBookmark := new(models.Bookmark)

	err := c.BodyParser(newBookmark)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := s.database.CreateBookmark(newBookmark.Name, newBookmark.Url)
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

//func (s *Service) UpdateBookmark(c *fiber.Ctx) error {
//	id := c.Params("id")
//
//	result, err := s.database.UpdateBookmark(id)
//	if err != nil {
//		return c.Status(503).SendString(err.Error())
//	}
//	return c.Status(200).JSON(result)
//
//}

func (s *Service) DeleteBookmark(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := s.database.DeleteBookmark(id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}
