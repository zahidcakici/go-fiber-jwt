package database

import (
	"fmt"
	"go-fiber-jwt-example/config"
	"go-fiber-jwt-example/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// Connect to the database
func Connect() {
	dsn := fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get("DB_PORT"), config.Get("DB_USER"), config.Get("DB_PASSWORD"), config.Get("DB_NAME"),
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
