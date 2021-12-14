package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
}

const DBConfig = "host=localhost user=user password=1234 dbname=user port=5432 sslmode=disable"

func InitDatabase() error {
	db, err := gorm.Open(postgres.Open(DBConfig), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Bookmark{})

	return nil
}

func CreateBookmark(b *Bookmark) (Bookmark, error) {
	db, err := gorm.Open(postgres.Open(DBConfig), &gorm.Config{})

	if err != nil {
		return *b, err
	}

	db.Create(b)

	return *b, nil
}

func GetBookmarks() ([]Bookmark, error) {
	bookmarks := []Bookmark{}
	db, err := gorm.Open(postgres.Open(DBConfig), &gorm.Config{})

	if err != nil {
		return bookmarks, err
	}

	db.Find(&bookmarks)

	return bookmarks, nil
}

func GetBookmark(id string) (Bookmark, error) {
	bookmark := Bookmark{}
	db, err := gorm.Open(postgres.Open(DBConfig), &gorm.Config{})

	if err != nil {
		return bookmark, err
	}

	db.Find(&bookmark, id)
	if bookmark.Name == "" {
		return bookmark, fmt.Errorf("can't find bookmark %s", id)
	}

	return bookmark, nil
}

func DeleteBookmark(id string) (Bookmark, error) {
	bookmark := Bookmark{}

	db, err := gorm.Open(postgres.Open(DBConfig), &gorm.Config{})
	if err != nil {
		return bookmark, err
	}

	db.First(&bookmark, id)
	if bookmark.Name == "" {
		return bookmark, fmt.Errorf("can't find bookmark %s", id)
	}

	db.Delete(&bookmark, id)

	return bookmark, nil
}
