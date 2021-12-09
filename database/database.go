package database

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type Bookmark struct {
  gorm.Model
  Name string `json:"name"`
  Url string `json:"url"`
}

func initDatabase() error {
  dsn := "host=localhost user=postgres password=1234 dbname=testdb port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return err
  }

  db.AutoMigrate()

  return nil
}
