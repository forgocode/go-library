package mail

import (
	"strings"

	"gopkg.in/gomail.v2"
)

type Options struct {
	Host     string
	Port     int
	User     string // 发件人
	Password string // 发件人密码
	MailTo   string // 收件人 多个用,分割
	Subject  string // 邮件主题
	Body     string // 邮件内容
}

func SendMail(opt *Options) error {

	m := gomail.NewMessage()

	//设置发件人
	m.SetHeader("From", opt.User)

	//设置发送给多个用户
	mailArrTo := strings.Split(opt.MailTo, ",")
	m.SetHeader("To", mailArrTo...)

	//设置邮件主题
	m.SetHeader("Subject", opt.Subject)

	//设置邮件正文
	m.SetBody("text/html", opt.Body)

	d := gomail.NewDialer(opt.Host, opt.Port, opt.User, opt.Password)

	return d.DialAndSend(m)
}
