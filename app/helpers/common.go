package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}

	return os.Getenv(key)
}

func CheckRequiredEnvs() {
	requiredEnvVars := []string{"SERVER_PORT", "DB_PORT", "DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "LOG_LEVEL"}

	for _, envVar := range requiredEnvVars {
		if value, exists := os.LookupEnv(envVar); !exists || value == "" {
			log.Panic(fmt.Sprintf("Error: Environment variable %v is not set.", envVar))
		}
	}
}
