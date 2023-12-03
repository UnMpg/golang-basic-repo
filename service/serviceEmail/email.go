package serviceemail

import (
	"crypto/tls"
	"go-project/config"
	"go-project/models"
	"go-project/utils/email"
	"go-project/utils/log"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

func SendEmailRegister(userData models.User, data models.EmailData) error {
	// bodyEmail := ""

	dataEmail := models.EmailStruck{
		From:    "",
		To:      userData.Email,
		Subject: data.Subject,
		Body:    email.BodySendEmailVerification(data.URL, "namssm"),
	}

	if err := SendEmail(dataEmail); err != nil {
		return err
	}

	return nil
}

func SendEmail(email models.EmailStruck) error {
	from := config.AppEnv.EmailFrom
	smtpHost := config.AppEnv.SMTPHost
	smtpUser := config.AppEnv.SMTPUser
	smtpPass := config.AppEnv.SMTPPass
	smtpPort := config.AppEnv.SMTPPort

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Body)
	m.AddAlternative("text/plain", html2text.HTML2Text(email.Body))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Log.Error("Error To send Email Dial")
		return err
	}

	return nil

}
