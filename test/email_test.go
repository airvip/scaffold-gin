package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "阿尔维奇 <sdqhwzb@163.com>"
	e.To = []string{"980062449@qq.com"}
	// e.Bcc = []string{"test_bcc@example.com"}
	// e.Cc = []string{"test_cc@example.com"}
	e.Subject = "邮件测试"
	e.Text = []byte("我是纯纯的文本信息")
	e.HTML = []byte("<h1>我是html文本信息</h1>")
	// err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
	// 返回 EOF 时，关闭 SSL
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "sdqhwzb@163.com", "我的登录密码，现在最好使用随机密码", "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}