package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ConfigUseSQLite      bool
	ConfigDebugDatabase  bool
	ConfigJWTSecret      []byte
	ConfigOptipngEnabled bool
)

func loadConfig() {
	var str string
	var is_present bool

	str, is_present = os.LookupEnv("USE_SQLITE")
	if is_present {
		ConfigUseSQLite = str == "true"
	} else {
		ConfigUseSQLite = false
	}

	str, is_present = os.LookupEnv("DEBUG_DATABASE")
	if is_present {
		ConfigDebugDatabase = str == "true"
	} else {
		ConfigDebugDatabase = false
	}

	str, is_present = os.LookupEnv("JWT_SECRET")
	if is_present {
		ConfigJWTSecret = []byte(str)
	} else {
		log.Panicln("No JWT secret configured!")
	}

	str, is_present = os.LookupEnv("ENABLE_OPTIPNG")
	if is_present {
		ConfigOptipngEnabled = str == "true"
	} else {
		ConfigOptipngEnabled = true
	}
}

func init() {
	log.Println("Loading config...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	loadConfig()
}
