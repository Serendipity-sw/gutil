// sftp帮助类
// create by gloomy 2017-05-09 20:19:24
package gutil

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// sftp配置类
// create by gloomy 2017-05-09 20:21:07
type SftpConfigStruct struct {
	Account      string // 登录用户名
	Password     string // 登录密码
	Port         int    // 服务器端口
	ConntionSize int    // MaxPacket sets the maximum size of the payload
}

// 建立sftp连接
func sftpConntion(sftpModel SftpConfigStruct, sftpClient *sftp.Client, sshClient *ssh.Client) {

}
