package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
}

const DBConfig = "host=localhost user=user password=1234 dbname=testdb port=5432 sslmode=disable"

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

func GetAllBookmarks() ([]Bookmark, error) {
	bookmarks := []Bookmark{}
	db, err := gorm.Open(postgres.Open(DBConfig), &gorm.Config{})

	if err != nil {
		return bookmarks, err
	}

	db.Find(&bookmarks)

	return bookmarks, nil
}
