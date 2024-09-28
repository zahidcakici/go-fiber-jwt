package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Get func to get env value
func Get(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
