package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get loads the .env file first, if possible. And then returns the value of
// the environment variable. Or an empty string, if there is no variable with a
// given key.
func Get(key string) string {
	// Try to load .env file.
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Failed to load .env file :/")
	}

	// Return the value.
	return os.Getenv(key)
}
