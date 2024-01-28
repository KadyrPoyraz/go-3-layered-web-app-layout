package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DBConfig
	App AppConfig
}

type DBConfig struct {
	Name     string
	Password string
	User     string
	Host     string
	Port     string
}

type AppConfig struct {
	Port string
}

func New() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbUser := os.Getenv("POSTGRES_USER")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	appPort := os.Getenv("APP_PORT")

	dbConfig := DBConfig{
		Name:     dbName,
		Password: dbPassword,
		User:     dbUser,
		Host:     dbHost,
		Port:     dbPort,
	}

	appConfig := AppConfig{
		Port: appPort,
	}

	cnf := Config{
		DB:  dbConfig,
		App: appConfig,
	}

	return cnf, nil
}
