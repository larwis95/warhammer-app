package db

import (
	"api/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
  dbName := os.Getenv("DATABASE_URL")
  db, err := gorm.Open(postgres.Open(dbName), &gorm.Config{TranslateError: true})
  if err != nil {
    panic("Failed to connect to database!")
  }
  return db
}

func AutoMigrate(db *gorm.DB) {
  db.AutoMigrate(&models.User{})
  db.AutoMigrate(&models.Allegiance{})
  db.AutoMigrate(&models.GrandAlliance{})
  db.AutoMigrate(&models.Unit{})
  db.AutoMigrate(&models.Ability{})
  db.AutoMigrate(&models.Weapon{})
  db.AutoMigrate(&models.DamageTable{})
}

