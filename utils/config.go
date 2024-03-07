package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ConfigDebugDatabase  bool
	ConfigJWTSecret      []byte
	ConfigOptipngEnabled bool
)

func loadConfig() {
	var str string
	var isPresent bool

	str, isPresent = os.LookupEnv("DEBUG_DATABASE")
	if isPresent {
		ConfigDebugDatabase = str == "true"
	} else {
		ConfigDebugDatabase = false
	}

	str, isPresent = os.LookupEnv("JWT_SECRET")
	if isPresent {
		ConfigJWTSecret = []byte(str)
	} else {
		log.Panicln("No JWT secret configured!")
	}

	str, isPresent = os.LookupEnv("ENABLE_OPTIPNG")
	if isPresent {
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
