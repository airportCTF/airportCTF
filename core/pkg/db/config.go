package db

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (c Config) ConnectionUrl() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Database)
}

func NewConfig(host string, port int, user, password, database string) Config {
	return Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
}

func NewConfigFromEnv() *Config {
	return &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     5432,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}
