package routes

import "os"

var secretKey = os.Getenv("JWT_SECRET")
