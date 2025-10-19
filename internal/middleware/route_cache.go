package middleware

import (
	"portfolioed/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func RouteCacheMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		entry, found := store.GetRouteBytes(c.Path())
		if !found {
			return c.Next()
		}
		log.Info().Str("Etag", entry.ETag).Str("path", c.Path()).Msg("Serving from route cache")
		c.Set("Content-Type", "text/html")
		c.Set("ETag", entry.ETag)
		if match := c.Get("If-None-Match"); match == entry.ETag {
			return c.SendStatus(fiber.StatusNotModified)
		}

		// This is correctly placed to only fire on a 200 OK, not a 304
		// capi.SendPageViewEvent(c)
		return c.Send(entry.Data)
	}
}
