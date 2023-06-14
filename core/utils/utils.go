package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	configFile "go-microservices.org/core/config"
	connection "go-microservices.org/core/connection"

	"github.com/jinzhu/copier"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"
	cache "github.com/patrickmn/go-cache"
)

const (
	// DateLayout ...
	DateLayout = "2006-01-02"
	// DatetimeLayout ...
	DatetimeLayout = "2006-01-02 15:04:05"
	// CoreUtils ...
	CoreUtils = "core.utils"
	// Postgres ...
	Postgres = "postgresql"
)

// CopyObject ...
func CopyObject(from interface{}, to interface{}) error {
	if from == nil {
		to = nil
		return nil
	}

	return copier.Copy(to, from)
}

// DatetimeToString ...
func DatetimeToString(input string) string {
	if len(input) >= 19 {
		return input[0:10] + " " + input[11:19]
	}
	return input
}

// StringToBool ...
func StringToBool(input string) (bool, error) {
	result, _ := strconv.ParseBool(input)
	return result, nil
}

// StringToInt ...
func StringToInt(input string) (int, error) {
	result, _ := strconv.Atoi(input)
	return result, nil
}

// StringToInt32 ...
func StringToInt32(input string) (int32, error) {
	result, _ := strconv.Atoi(input)
	return int32(result), nil
}

//StringToFloat32 ...
func StringToFloat32(input string) (float32, error) {
	result, _ := strconv.ParseFloat(input, 32)
	return float32(result), nil
}

//StringToFloat64 ...
func StringToFloat64(input string) (float64, error) {
	result, _ := strconv.ParseFloat(input, 64)
	return result, nil
}

// StringToJSONMap ...
func StringToJSONMap(input string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(input), &result)
	return result, err
}

// StringToDuration ...
func StringToDuration(input string) (time.Duration, error) {
	result, _ := strconv.Atoi(input)
	return time.Duration(result), nil
}

// UniqueValue ... to remove duplicate on array of string supplied
func UniqueValue(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// SendLogError ...
func SendLogError(id string, err error) error {
	log.Log(errors.BadRequest(id, err.Error()))
	return err
}

// SendLogInfo ...
func SendLogInfo(msg ...string) {
	log.Info(msg)
}

var (
	conn    connection.Connection
	oneConn sync.Once
)

// GetPostgresHandler ...
func GetPostgresHandler() connection.Connection {
	oneConn.Do(func() {
		conn = connection.NewPostgresConnection()
	})

	return conn
}

// GetDatasourceInfo ...
func GetDatasourceInfo() string {
	if typ := os.Getenv("REPO_TYPE"); typ != "" {
		return typ
	}

	return Postgres
}

var (
	caches    *cache.Cache
	onceCache sync.Once
)

// GetCacheHandler ...
func GetCacheHandler() *cache.Cache {
	config := configFile.GetApplicationConfig()
	onceCache.Do(func() {
		caches = cache.New(config.CacheExpiry*time.Second, config.CacheCleanup*time.Second)
	})

	return caches
}

// GetCache ...
func GetCache(key string, cache *cache.Cache) interface{} {
	result, found := cache.Get(key)
	if found {
		return result
	}
	return nil
}

// SetCache ...
func SetCache(key string, value interface{}, cacheParam *cache.Cache) interface{} {
	cacheParam.Set(key, value, cache.DefaultExpiration)
	return value
}

// GetCacheKey ...
func GetCacheKey(params ...string) string {
	var fnName string
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		fnName = "?"
	} else {
		fn := runtime.FuncForPC(pc)
		fnName = fn.Name()
	}

	return fnName + "_" + strings.Join(params, "|")
}

// FindString ...
func FindString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// NVL checks if the argument is zero or null
func NVL(obj interface{}, replc interface{}) interface{} {

	if obj == nil {
		if replc == nil {
			return nil
		}
		return replc
	}

	return obj

}

// GetHashPassword ...
func GetHashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}

// GetLimitOffset ...
func GetLimitOffset(page, limit int) (int, int) {
	if limit <= 0 {
		limit = configFile.GetApplicationConfig().DefaultPageLimit
	}

	var offset int
	if offset = (page - 1) * limit; offset <= 0 {
		offset = 0
	}

	return limit, offset
}

// TrimLower ...
func TrimLower(val string) string {
	return strings.ToLower(strings.TrimSpace(val))
}

// TrimUpper ...
func TrimUpper(val string) string {
	return strings.ToUpper(strings.TrimSpace(val))
}
