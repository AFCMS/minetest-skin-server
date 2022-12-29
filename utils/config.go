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
	ConfigUseSQLite      bool
	ConfigDebugDatabase  bool
	ConfigJWTSecret      []byte
	ConfigOptipngEnabled bool
)

func loadConfig() {
	ConfigUseSQLite = rEnv("USE_SQLITE") == "true"

	ConfigDebugDatabase = rEnv("DEBUG_DATABASE") == "true"

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
