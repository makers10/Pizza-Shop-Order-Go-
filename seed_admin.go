package main

import (
	"log"
	"pizza-tracker-go/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=Harsh@123 dbname=pizza_shop port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Generate bcrypt hash for password "admin123"
	hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	user := models.User{
		Username: "admin",
		Password: string(hash),
	}

	// Create user (ignore error if it already exists)
	result := db.Where(models.User{Username: "admin"}).FirstOrCreate(&user)
	if result.Error != nil {
		log.Println("Note:", result.Error)
	} else {
		log.Println("Admin user 'admin' with password 'admin123' ensured in DB.")
	}
}
