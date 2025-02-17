package pkg

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
) 

type pkgisntance struct{}
type Messenger interface {
	Mail_sender(Subject string ,msg string,REmail string) bool
}

func GetpkgIntstance() Messenger {
	return &pkgisntance{}
}

func (P *pkgisntance) Mail_sender(Subject string ,msg string,REmail string) bool {
	smtpHost :=  os.Getenv("ORG_SMTP_HOST")
	smtpPort :=  os.Getenv("ORG_SMTP_PORT")
	senderEmail := os.Getenv("ORG_EMAIL")
	senderPassword := os.Getenv("ORG_APP_PASS")
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", REmail)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/plain", msg)
	Port,_:=strconv.Atoi(smtpPort)
	d := gomail.NewDialer(smtpHost, Port, senderEmail, senderPassword)

	if err := d.DialAndSend(m); err != nil {
		panic("error is here")
		return false
	}
	return true

}