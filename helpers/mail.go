package helpers

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(mail string, code uint) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "Todo@gmail.com")

	m.SetHeader("To", mail)

	m.SetHeader("Subject", "Şifre sıfırlama")

	m.SetBody("text/plain", "code : "+strconv.Itoa(int(code)))

	d := gomail.NewDialer("smtp.gmail.com", 587, "2194257061@nku.edu.tr", "***")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
