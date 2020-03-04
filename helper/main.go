package helper

import (
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload" // initialize
)

// GetUUID func
func GetUUID() string {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Println(err)
	}

	return uuid.String()
}

// GetEnv func
func GetEnv(name string) string {
	env := os.Getenv(name)

	if env == "" {
		log.Fatalf("Environment variable not found: \"%s\".", name)
	}

	return env
}

// IsLocal func
func IsLocal() bool {
	return GetEnv("APP_ENV") == "local"
}

// IsProduction func
func IsProduction() bool {
	return GetEnv("APP_ENV") == "production"
}
