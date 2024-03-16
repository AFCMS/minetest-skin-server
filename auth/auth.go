package auth

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory/v2"
)

var SessionStore *session.Store

func Initialize() {
	SessionStore = session.New(session.Config{
		Expiration:        24 * time.Hour,
		Storage:           memory.New(), // TODO: use Redis
		KeyLookup:         "cookie:session_id",
		CookieDomain:      "",
		CookiePath:        "",
		CookieSecure:      false, // TODO: handle production
		CookieHTTPOnly:    true,
		CookieSessionOnly: false,
	})
}
