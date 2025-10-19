package util

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"portfolioed/internal/store"
	"sync"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

var bufferPool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)
	err := component.Render(c.Context(), buf)
	if err != nil {
		return err
	}

	renderedHTML := make([]byte, buf.Len())
	copy(renderedHTML, buf.Bytes())
	go store.SetRouteBytes(c.Path(), renderedHTML)

	h := fnv.New64a()
	h.Write(renderedHTML)

	etag := fmt.Sprintf("\"%x\"", h.Sum64())
	c.Set("ETag", etag)

	if match := c.Get("If-None-Match"); match == etag {
		return c.SendStatus(fiber.StatusNotModified)
	}
	//TODO: Add it back When we are Deploying
	// capi.SendPageViewEvent(c)
	return c.Send(renderedHTML)
}
