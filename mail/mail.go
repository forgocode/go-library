package mail

import (
	"fmt"
	"net/smtp"
)

type Email struct {
	From     string
	To       []string
	Body     string
	AuthInfo smtp.Auth
	Setting  *EmailSetting
}

type EmailSetting struct {
	UserName string
	Password string
	Host     string
	Port     string
}

func NewEmailSetting() *EmailSetting {
	return new(EmailSetting)
}

func NewEmail() *Email {
	return new(Email)
}

func Auth(userName, password, host, port string) smtp.Auth {

	return smtp.PlainAuth("", userName, password, host)

}

func SendEmail(from, msg string, to []string, auth smtp.Auth, setting *EmailSetting) error {
	m := NewEmail()
	m.Setting = NewEmailSetting()
	m.Setting = setting
	m.AuthInfo = auth
	m.From = from
	m.To = to
	m.Body = msg
	return m.sendMail()
}

func GetEmailSetting(userName, password, host, port string) *EmailSetting {
	s := NewEmailSetting()
	s.UserName = userName
	s.Password = password
	s.Host = host
	s.Port = port
	return s
}

func (e *Email) sendMail() error {
	e.AuthInfo = smtp.PlainAuth("", e.From, e.Setting.Password, e.Setting.Host)
	fmt.Printf("11111%+v\n", e)
	return smtp.SendMail(e.Setting.Host+":"+e.Setting.Port, e.AuthInfo, e.From, e.To, []byte(e.Body))
}
