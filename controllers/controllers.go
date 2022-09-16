package controllers

import (
	"bookmark-api-fiber/services"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) SetUpRoutes(app *fiber.App) {

	app.Get("/", c.service.Status)

	app.Get("/api/bookmark", c.service.GetAllBookmarks)
	app.Get("/api/bookmark/:id", c.service.GetBookmarkId)
	app.Post("/api/bookmark", c.service.SaveBookmark)
	//app.Put("/api/bookmark/:id", c.service.UpdateBookmark)
	app.Delete("/api/bookmark/:id", c.service.DeleteBookmark)

}
