package store

import (
	"fmt"
	"hash/fnv"

	"github.com/rs/zerolog/log"
)

type CacheEntry struct {
	Data []byte
	ETag string
}

func SetRouteBytes(path string, data []byte) {
	h := fnv.New64a()
	h.Write(data)
	etag := fmt.Sprintf("\"%x\"", h.Sum64())
	entry := CacheEntry{
		Data: data,
		ETag: etag,
	}
	log.Info().Str("path", path).Str("etag", etag).Msg("caching route")

	s := Get()
	s.Set(path, entry)
}

func GetRouteBytes(path string) (CacheEntry, bool) {
	s := Get()
	v, keyExists := s.Get(path)
	if !keyExists {
		return CacheEntry{}, false
	}
	u, typeOK := v.(CacheEntry)
	return u, typeOK
}
