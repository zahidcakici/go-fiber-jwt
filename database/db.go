package database

import (
	"fmt"
	"go-fiber-jwt-example/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// Connect to the database
func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	// Migrate the schemas
	err = DB.AutoMigrate(&model.Book{}, &model.User{})
	if err != nil {
		panic("failed to migrate database schemas")
	}
	fmt.Println("Database Migrated")
}
