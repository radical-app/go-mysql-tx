package mysql

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const DEFAULT_CNN_FLAGS = "parseTime=true&multiStatements=true&loc=UTC&charset=utf8"

type ConfigOpenable interface {
	GetConnection() string
}

type Config struct{
	User              string
	Password          string
	DbName            string
	Host              string
	Port              string

	CnnFlags           string
}


func ConfigFromEnvs(prefix string) *Config {
	return ConfigFromEnvsWithCNNFlags(prefix, DEFAULT_CNN_FLAGS)
}


func ConfigFromEnvsWithCNNFlags(prefix string, flags string) *Config {
	if prefix == "" {
		prefix = "BOOKING"
	}
	prefix = strings.ToUpper(strings.TrimRight(prefix, "_")+"_")

	_ = godotenv.Load("./.env" )

	return &Config{
		os.Getenv(prefix + "DB_USER"),
		os.Getenv(prefix + "DB_PASSWORD"),
		os.Getenv(prefix + "DB_NAME"),
		os.Getenv(prefix + "DB_HOST"),
		os.Getenv(prefix + "DB_PORT"),
		flags,
	}
}


func (c *Config) GetConnection() string {
	return c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DbName + "?"+c.CnnFlags
}
