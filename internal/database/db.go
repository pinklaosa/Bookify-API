package database

import (
	"Bookify/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{}) 
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Database migrated successfully!")
}


func ConnectDb() *gorm.DB{
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

	// MigrateDB(db)
	DB = db
	return DB
}


