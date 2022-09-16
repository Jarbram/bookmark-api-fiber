package main

import (
	"bookmark-api-fiber/controllers"
	"bookmark-api-fiber/database"
	"bookmark-api-fiber/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		return
	}

	database := database.NewDatabase(db)
	service := services.NewService(database)
	controller := controllers.NewController(service)

	controller.SetUpRoutes(app)
	app.Listen(":3000")
}
