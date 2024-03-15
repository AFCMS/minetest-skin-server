package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// ConfigFrontendDevMode is true if the frontend is in development mode (served externally and proxied by the backend)
	ConfigFrontendDevMode bool
	// ConfigFrontendURL is the URL of the frontend when in development mode
	ConfigFrontendURL    string
	ConfigDebugDatabase  bool
	ConfigOptipngEnabled bool
)

func loadConfig() {
	var str string
	var isPresent bool

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_FRONTEND_DEV_MODE")
	if isPresent {
		ConfigFrontendDevMode = str == "true"
	} else {
		ConfigFrontendDevMode = false
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_FRONTEND_URL")
	if isPresent && ConfigFrontendDevMode {
		ConfigFrontendURL = str
	} else {
		ConfigFrontendURL = ""
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_DATABASE_LOGGING")
	if isPresent {
		ConfigDebugDatabase = str == "true"
	} else {
		ConfigDebugDatabase = false
	}

	str, isPresent = os.LookupEnv("MT_SKIN_SERVER_ENABLE_OPTIPNG")
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
