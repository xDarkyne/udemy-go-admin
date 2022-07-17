package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	EnvDevelopment string = "development"
	EnvProduction  string = "production"
)

type AppConfig struct {
	DB             DatabaseConfig
	Environment    string `env:"APP_ENV" env-default:"development"`
	Port           int    `env:"APP_PORT" env-default:"3000"`
	TimeZone       string `env:"APP_TZ" env-default:"Europe/Berlin"`
	AuthCookieName string `env:"AUTH_COOKIE_NAME" env-default:"session"`
	AuthSecret     []byte `env:"AUTH_SECRET" env-default:"secret"`
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Username string `env:"DB_USER" env-default:"dev"`
	Password string `env:"DB_PASS" env-default:"dev"`
	DBName   string `env:"DB_NAME" env-default:"dev"`
	Port     int    `env:"DB_PORT" env-default:"5432"`
	SSLMode  string `env:"DB_SSL_MODE" env-default:"disable"`
}

var App AppConfig

func LoadConfig() {
	var err error
	preEnv := os.Getenv("APP_ENV")

	if preEnv != EnvProduction {
		fmt.Println("Loading environment from .env")
		err = godotenv.Load(".env")

		if err != nil {
			fmt.Println(err)
		}
	}

	err = cleanenv.ReadEnv(&App)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
