package email

import (
	"errors"
	"net/smtp"
)

type Emailer struct {
	host     string
	port     string
	username string
	password string
}

func NewEmailer(host, port, username, password string) *Emailer {
	return &Emailer{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (e *Emailer) SendEmail(to, subject, body string) error {
	from := e.username
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(e.host+":"+e.port,
		smtp.PlainAuth("", e.username, e.password, e.host),
		from, []string{to}, []byte(msg))

	if err != nil {
		return errors.New("could not send the email")
	}

	return nil
}
