package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func rEnv(key string) string {
	return os.Getenv(key)
}

var (
	ConfigJWTSecret      []byte
	ConfigOptipngEnabled bool
)

func loadConfig() {
	ConfigJWTSecret = []byte(rEnv("JWT_SECRET"))

	ConfigOptipngEnabled = rEnv("ENABLE_OPTIPNG") == "true"
}

func init() {
	log.Println("Loading config...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	loadConfig()
}
