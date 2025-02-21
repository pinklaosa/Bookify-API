package database

import (
	"Bookify/internal/config"
	"Bookify/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func MigrateDB(db *gorm.DB) {
	// err := db.AutoMigrate(&models.Book{}) 
	err := db.AutoMigrate(&models.Category{}) 
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Database migrated successfully!")
}


func ConnectDb() *gorm.DB{
	config.LoadConfig()
	cfg := config.AppConfigInstance.Database;

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

	// MigrateDB(db)
	DB = db
	return DB
}


