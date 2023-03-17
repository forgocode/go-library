package mail

import (
	"fmt"
	"testing"
)

func Test_SendMail(t *testing.T) {
	opt := &Options{
		Host:     "smtp.163.com",
		Port:     465,
		User:     "GgoCoderAdmin@163.com",
		Password: "OWOVSUUHBKJHLIKR",
		MailTo:   "forgocode@163.com",
		Subject:  "test title",
		Body:     "test body",
	}
	fmt.Println(SendMail(opt))
	// a := Auth("GgoCoderAdmin", "OWOVSUUHBKJHLIKR", "smtp.163.com", "465")
	// s := GetEmailSetting("GgoCoderAdmin", "OWOVSUUHBKJHLIKR", "smtp.163.com", "465")
	// fmt.Println("111", SendEmail("GgoCoderAdmin@163.com", "test", []string{"GgoCoder@163.com"}, a, s))
}
