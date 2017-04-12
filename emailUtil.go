// email 类库
// create by gloomy 2017-04-12 11:16:21
package common

import (
	"net/smtp"
)

// 发送邮件
//
func SendEmail() {
	auth := smtp.PlainAuth("", "shaow@axon.com.cn", "Sw435464", "smtp.qiye.163.com")
	to := []string{"sw.gloomysw@gmail.com", "chenm@axon.com.cn"}
	msg := []byte("To: bain@axon.com.cn;shaow@axon.com.cn;chenm@axon.com.cn\r\n" +
		"Subject: 测试golang邮件发送 \r\n" +
		"Content-Type: text/plain; charset=UTF-8" + "\r\n\r\n")
	//for _, value := range sendMailContent {
	msg = append(msg, []byte("测试邮件发送")...)
	//}
	err := smtp.SendMail("smtp.qiye.163.com:25", auth, "shaow@axon.com.cn", to, msg)
	if err != nil {
		//glog.Error("邮件发送失败! err: %s \n", err.Error())
		//	fmt.Println(err.Error())
		return
	}
}
