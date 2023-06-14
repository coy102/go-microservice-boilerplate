package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	common "go-microservices.org/core/model"

	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"
)

type config struct {
	ApplicationConfig *common.ApplicationConfig
	FileModified      time.Time
}

var (
	appConfig *config
	once      sync.Once
)

func getApplicationConfigInstance() *config {
	once.Do(func() {
		appConfig = new(config)
	})

	return appConfig
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

// GetApplicationConfig ..
func GetApplicationConfig() *common.ApplicationConfig {
	fileName := basepath + "/.app.config.json"

	file, err := os.Stat(fileName)
	if err != nil {
		log.Log(errors.BadRequest("common.config.GetApplicationConfig", err.Error()))
		return nil
	}

	instance := getApplicationConfigInstance()
	if file.ModTime().After(instance.FileModified) {
		log.Info("Load application config from files ..")

		jsonFile, err := os.Open(fileName)
		if err != nil {
			log.Log(errors.BadRequest("common.config.GetApplicationConfig", err.Error()))
			return nil
		}

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var appConfig *common.ApplicationConfig

		json.Unmarshal([]byte(byteValue), &appConfig)

		instance.ApplicationConfig = appConfig
		instance.FileModified = file.ModTime()
	}

	return instance.ApplicationConfig
}
