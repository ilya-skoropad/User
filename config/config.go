package config

import "os"

type config struct {
	AppHost string
	AppPort string
}

func Get() config {
	return config{
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	}
}
