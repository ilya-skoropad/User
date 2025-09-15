package service

import (
	"ilya-skoropad/user/config"
	"net/smtp"
)

type Email interface {
	Send(to string, title string, body string) error
}

type emailSender struct {
	conf config.Config
}

func (es *emailSender) Send(to string, title string, body string) error {
	from := es.conf.AppHost
	password := es.conf.SmtpSender

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + title + "\n\n" +
		body

	err := smtp.SendMail(
		es.conf.SmtpProvider,
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg),
	)

	return err
}

func NewEmailSender(conf config.Config) Email {
	return &emailSender{}
}
