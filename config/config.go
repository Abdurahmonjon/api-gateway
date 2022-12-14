package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production

	StudentServiceHost string
	StudentServicePort int

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("EVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.StudentServiceHost = cast.ToString(getOrReturnDefault("STUDENT_SERVICE_HOST", "127.0.0.1"))
	c.StudentServicePort = cast.ToInt(getOrReturnDefault("student_SERVICE_PORT", 50051))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 100000000))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
