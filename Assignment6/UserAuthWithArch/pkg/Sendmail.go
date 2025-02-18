package pkg

import (
	"strconv"

	"github.com/TusharKM1224/config"
	"gopkg.in/gomail.v2"
)

type pkgisntance struct {
	configs config.Configs_data
}
type Messenger interface {
	Mail_sender(Subject string, msg string, REmail string) bool
}

func GetpkgIntstance(C config.Configs_data) Messenger {
	return &pkgisntance{configs: C}
}

func (P *pkgisntance) Mail_sender(Subject string, msg string, REmail string) bool {
	smtpHost := P.configs.Smtp_configs.Org_smtp_host
	smtpPort := P.configs.Smtp_configs.Org_smtp_port
	senderEmail := P.configs.Smtp_configs.Org_email
	senderPassword := P.configs.Smtp_configs.Org_App_Pass
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", REmail)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/plain", msg)
	Port, _ := strconv.Atoi(smtpPort)
	d := gomail.NewDialer(smtpHost, Port, senderEmail, senderPassword)

	if err := d.DialAndSend(m); err != nil {
		return false
	}
	return true

}
