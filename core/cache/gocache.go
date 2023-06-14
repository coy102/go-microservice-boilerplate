package cache

import (
	"os"
	"time"

	"go-microservices.org/core/utils"

	gcache "github.com/patrickmn/go-cache"
)

type goCache struct {
	cache *gcache.Cache
}

func setupGocache() ICache {
	expire, _ := utils.StringToDuration(os.Getenv("SESSION_EXPIRE_HOUR"))
	expireHour := expire * time.Hour
	session := &goCache{
		cache: gcache.New(expireHour, expireHour),
	}

	return session
}

func (ses *goCache) CreateKey(value string) string {
	return prefixKey + "_" + value
}

func (ses *goCache) GetKeyExist(key string) (interface{}, bool, error) {
	result, found := ses.cache.Get(key)
	if !found {
		return nil, false, nil
	}

	return result, true, nil
}

func (ses *goCache) SetKey(key string, value interface{}) error {
	ses.cache.SetDefault(key, value)
	return nil
}

func (ses *goCache) SetKeyWithExpire(key string, value interface{}, expire time.Duration) error {
	ses.cache.Set(key, value, expire)
	return nil
}

func (ses *goCache) DeleteKey(key string) error {
	ses.cache.Delete(key)
	return nil
}
