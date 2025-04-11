package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(environment string) {
	if environment == "development" {
		envDevelopmentPath := os.Getenv("ENV_DEVELOPMENT_PATH")

		if envDevelopmentPath == "" {
			envDevelopmentPath = ".env.development"
		}

		envPath, _ := filepath.Abs(envDevelopmentPath)

		err := godotenv.Load(envPath)
		if err != nil {
			panic("Error loading .env.development file")
		}
	}

	config = viper.New()
	config.SetConfigName("config")
	config.AutomaticEnv()
}

func GetEnv(key string) string {
	return config.GetString(key)
}

func GetEnvInt(key string) int {
	return config.GetInt(key)
}
