package mail

import (
	"fmt"
	"testing"
)

func Test_SendMail(t *testing.T) {
	a := Auth("GgoCoderAdmin", "OWOVSUUHBKJHLIKR", "smtp.163.com", "465")
	s := GetEmailSetting("GgoCoderAdmin", "OWOVSUUHBKJHLIKR", "smtp.163.com", "465")
	fmt.Println("111", SendEmail("GgoCoderAdmin@163.com", "test", []string{"GgoCoder@163.com"}, a, s))
}
