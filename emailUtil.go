// email 类库
// create by gloomy 2017-04-12 11:16:21
package gutil

import (
	"fmt"
	"net/smtp"
	"strings"
)

// 发送邮件
// create by gloomy 2017-04-12 11:18:23
func SendEmail(account, passWord, smtpUrl, smtpUrlPort, emailTitle string, emailContent *[]byte, toEmailUser []string) error {
	auth := smtp.PlainAuth("", account, passWord, smtpUrl)
	to := toEmailUser
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/plain; charset=UTF-8"+"\r\n\r\n", strings.Join(toEmailUser, ";"), emailTitle))
	msg = append(msg, (*emailContent)...)
	return smtp.SendMail(fmt.Sprintf("%s:%s", smtpUrl, smtpUrlPort), auth, account, to, msg)
}
