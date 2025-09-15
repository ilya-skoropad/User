package config

import "os"

type Config struct {
	AppHost string
	AppPort string
	DbCon   string

	DoNotSendEmails bool
	SmtpSender      string
	SmtpPass        string
	SmtpProvider    string
}

func Get() Config {
	return Config{
		AppHost:         os.Getenv("APP_HOST"),
		AppPort:         os.Getenv("APP_PORT"),
		DbCon:           os.Getenv("DB_CON"),
		DoNotSendEmails: os.Getenv("DO_NOT_SEND_EMAILS") != "",
		SmtpSender:      os.Getenv("SMTP_SENDER"),
		SmtpPass:        os.Getenv("SMTP_PASS"),
		SmtpProvider:    os.Getenv("SMTP_PROVIDER"),
	}
}
