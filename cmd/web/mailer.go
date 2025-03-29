package main

import (
	gomail "gopkg.in/mail.v2"
)

type Mailer struct {
	client *gomail.Dialer
	sender string
}

func newMailer(host string, port int, username string, password string) (*Mailer, error) {
	dialer := gomail.NewDialer(host, port, username, password)

	return &Mailer{
		client: dialer,
		sender: username,
	}, nil
}

func (m *Mailer) Send(clientEmail, clientSubject, clientMessage string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", clientEmail)
	message.SetHeader("To", m.sender)
	message.SetHeader("Subject", clientSubject)

	message.SetBody("text/plain", clientMessage+" - from: "+clientEmail)

	if err := m.client.DialAndSend(message); err != nil {
		return err
	} else {
		return nil
	}
}
