package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func NewConfig() (*Config, error) {
	var Config Config
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error can't get the environment variables by file")
		os.Exit(1)
	}
	if err := env.Parse(&Config); err != nil {
		logrus.Fatalf("Error initializing: %s", err.Error())
		os.Exit(1)
	}
	return &Config, nil
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Bogota", c.DB.HOST, c.DB.USER, c.DB.PASSWORD, c.DB.DBNAME)
}

type Config struct {
	APP
	DB
	Auth
}

type Auth struct {
	Secret        string `env:"AUTH_SECRET" envDefault:"ionix"`
	TokenDuration int    `env:"AUTH_TOKEN_DURATION" envDefault:"4"`
}

type DB struct {
	HOST     string `env:"DB_HOST" envDefault:"localhost"`
	USER     string `env:"DB_USER" envDefault:"bia"`
	PASSWORD string `env:"DB_PASSWORD" envDefault:""`
	DBNAME   string `env:"DB_NAME" envDefault:"power_consumption"`
	PORT     string `env:"DB_PORT" envDefault:"3306"`
	SSLMODE  string `env:"DB_SSL_MODE" envDefault:"disable"`
	TIMEZONE string `env:"DB_TIME_ZONE"  envDefault:"America/Bogota"`
}

type APP struct {
	PORT string `env:"APP_PORT" envDefault:"8080"`
}
