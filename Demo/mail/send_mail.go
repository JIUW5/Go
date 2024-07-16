package mail

import (
	"log"
	"net/smtp"
)

const (
	username = "2923848605@qq.com"
	password = "自己的邮箱密码"
	host     = "smtp.qq.com"
	addr     = "smtp.qq.com:25"
)

func SendMail() {
	// PlainAuth函数的第一个参数是identity，这个参数一般为空字符串即可,第二个参数是用户名，第三个参数是密码，第四个参数是主机名
	auth := smtp.PlainAuth("", username, password, host)
	user := "2923848605@qq.com"
	to := []string{"jiuw596@gmail.com"}
	//msg是一个字节切片，包含了邮件的内容
	msg := []byte(`From: 2923848605@qq.com
To: jiuw596@gmail.com
Subject: 测试邮件

这是邮件内容
`)

	//SendMail函数的第一个参数是邮件服务器的地址，第二个参数是认证信息，第三个参数是发件人的邮箱，第四个参数是收件人的邮箱，第五个参数是邮件的内容
	err := smtp.SendMail(addr, auth, user, to, msg)

	if err != nil {
		log.Printf("Error when send mail: %s", err.Error())
	}
}
