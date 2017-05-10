// sftp帮助类
// create by gloomy 2017-05-09 20:19:24
package gutil

import (
	"errors"
	"github.com/pkg/sftp"
	"github.com/swgloomy/crypto/ssh"
	"github.com/swgloomy/crypto/ssh/agent"
	"net"
	"os"
)

// sftp配置类
// create by gloomy 2017-05-09 20:21:07
type SftpConfigStruct struct {
	Account      string // 登录用户名
	Password     string // 登录密码
	Port         int    // 服务器端口
	ConntionSize int    // MaxPacket sets the maximum size of the payload
	Addr         string // 连接地址
}

// 建立sftp连接
// create by gloomy 2017-05-10 11:32:17
func sftpConntion(sftpModel SftpConfigStruct, sftpClient *sftp.Client, sshClient *ssh.Client) error {
	var (
		auths []ssh.AuthMethod
		err   error
	)
	if aconn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(aconn).Signers))
	}
	auths = append(auths, ssh.Password(sftpModel.Password))
	configModel := ssh.ClientConfig{
		User:            sftpModel.Account,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshClient, err = ssh.Dial("tcp", sftpModel.Addr, &configModel)
	if err != nil {
		return err
	}
	sftpClient, err = sftp.NewClient(sshClient, sftp.MaxPacket(If(sftpModel.ConntionSize > 0, sftpModel.ConntionSize, 1<<15).(int)))

	return err
}

// sftp 关闭
// create by gloomy 2017-05-10 12:06:05
func SftpClose(sftpClient *sftp.Client, sshClient *ssh.Client) error {
	var (
		errs     string
		boNumber = 0
	)
	err := sftpClient.Close()
	if err != nil {
		errs += errors.New(err.Error())
		boNumber++
	}
	err = sshClient.Close()
	if err != nil {
		errs += errors.New(err.Error())
		return errors.New(errs)
	}
	if boNumber > 0 {
		return errors.New(errs)
	}
	return err
}

// sftp读取文件夹内容
// create by gloomy 2017-05-10 11:50:14
func SftpReadDir(sftpModel SftpConfigStruct, sftpClient *sftp.Client, sshClient *ssh.Client, pathStr string) (*[]os.FileInfo, error) {
	if sftpClient == nil || sshClient == nil {
		err := sftpConntion(sftpModel, sftpClient, sshClient)
		if err != nil {
			return nil, err
		}
	}
	files, err := sftpClient.ReadDir(pathStr)
	if err != nil {
		err := sftpConntion(sftpModel, sftpClient, sshClient)
		if err != nil {
			return nil, err
		}
		files, err = sftpClient.ReadDir(pathStr)
		if err != nil {
			return nil, err
		}
	}
	return &files, err
}
