package cache

import (
	"os"
	"time"
)

const (
	gocache   = "gocache"
	redis     = "redis"
	prefixKey = "_ses"
)

// ICache ...
type ICache interface {
	CreateKey(value string) string
	GetKeyExist(key string) (interface{}, bool, error)
	SetKey(key string, value interface{}) error
	SetKeyWithExpire(key string, value interface{}, expire time.Duration) error
	DeleteKey(key string) error
}

var mapSetupCache map[string]func() ICache

func init() {
	mapSetupCache = make(map[string]func() ICache)
	mapSetupCache[gocache] = setupGocache
}

// NewSession ...
func NewSession() ICache {
	typ := os.Getenv("SESSION_TYPE")
	if f, ok := mapSetupCache[typ]; ok {
		return f()
	}

	return setupGocache()
}
