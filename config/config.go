package config

import "os"

type Config struct {
	AppHost string
	AppPort string
	DbCon   string
}

func Get() Config {
	return Config{
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),
		DbCon:   os.Getenv("DB_CON"),
	}
}
