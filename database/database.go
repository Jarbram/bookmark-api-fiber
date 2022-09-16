package database

import (
	"bookmark-api-fiber/models"

	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase(c *gorm.DB) *Database {
	return &Database{Client: c}
}

func (d *Database) GetAllBookmarks() ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark

	d.Client.Find(&bookmarks)

	return bookmarks, nil
}

func (d *Database) GetBookmarkId(id string) (*models.Bookmark, error) {

	var bookmarks models.Bookmark

	d.Client.Find(&bookmarks, id)

	return &bookmarks, nil
}

func (d *Database) CreateBookmark(name string, url string) (models.Bookmark, error) {
	var newBookmark = models.Bookmark{Name: name, Url: url}

	d.Client.Create(&models.Bookmark{Name: name, Url: url})

	return newBookmark, nil
}

//func (d *Database) UpdateBookmark(id string) (*models.Bookmark, error) {
//	bookmark := new(models.Bookmark)
//
//	d.Client.Where("id = ?", id).Updates(bookmark)
//	return bookmark, nil
//}

func (d *Database) DeleteBookmark(id string) (*models.Bookmark, error) {

	var bookmarks models.Bookmark

	d.Client.Delete(&bookmarks, id)

	return &bookmarks, nil
}
