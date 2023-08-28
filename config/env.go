package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"correspondence-composer/utils/log"
)

func LoadEnvFile(logger log.Logger) {
	path, _ := os.Getwd()
	for {
		envFile := path + "/.env"
		if _, err := os.Stat(envFile); err == nil {
			err := godotenv.Load(envFile)
			if err != nil {
				logger.Warn(err)
			}
		}
		if len(path) <= 1 {
			break
		}

		path = filepath.Dir(path)
	}
}

func getEnvOrDefault(name, defaultValue string) string {
	value, found := os.LookupEnv(name)
	if !found {
		value = defaultValue
	}
	return value
}
