package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env.development")
	if err != nil {
		fmt.Print("Error loading .env.development file")
	}
	// Return the value of the variable
	return os.Getenv(key)
}
