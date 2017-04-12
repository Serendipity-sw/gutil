// FTP工具类
// create by gloomy 2017-4-1 17:32:06
package common

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

// FTP帮助类实体
// create by gloomy 2017-4-1 17:34:16
type FtpHelpStruct struct {
	IpAddr    string        // ip 地址
	Port      int           // 端口
	TimeOut   time.Duration // 超时时间
	UserName  string        // 用户名
	PassWord  string        // 密码
	FilePaths string        // 目标服务器路径
}

var (
	ftpConntion *ftp.ServerConn
)

// FTP文件传输
// create by gloomy 2017-4-1 17:36:11
// FTP配置实体 文件内容 创建目标服务器的文件名
// 错误对象
func FtpFileStor(model *FtpHelpStruct, contentByte *[]byte, createFilePath string) error {
	var (
		err error
	)
	if ftpConntion == nil || ftpConntion.NoOp() != nil {
		ftpConntion, err = ftpLogin(model)
		if err != nil {
			return err
		}
	}
	return ftpConntion.Stor(createFilePath, bytes.NewReader(*contentByte))
}

// FTP登录
// create by gloomy 2017-4-1 17:39:59
// 输入参数 FTP配置实体
// 输出参数 FTP连接对象 错误对象
func ftpLogin(model *FtpHelpStruct) (*ftp.ServerConn, error) {
	c, err := ftp.DialTimeout(fmt.Sprintf("%s:%d", model.IpAddr, model.Port), model.TimeOut)
	if err != nil {
		return nil, err
	}
	err = c.Login(model.UserName, model.PassWord)
	if err != nil {
		return nil, err
	}
	err = c.NoOp()
	if err != nil {
		return nil, err
	}
	return c, err
}

// FTP文件删除
// create by gloomy 2017-04-02 01:08:15
// 文件名 ftp配置对象
// 错误对象
func FtpRemoveFile(filePathStr string, model *FtpHelpStruct) error {
	var (
		err error
	)
	if ftpConntion == nil || ftpConntion.NoOp() != nil {
		ftpConntion, err = ftpLogin(model)
		if err != nil {
			return err
		}
	}
	return ftpConntion.Delete(filePathStr)
}

// ftp修正远程服务器文件名称
// create by gloomy 2017-04-04 21:26:48
// 源文件 修正后的文件名称 ftp配置对象
// 错误对象
func FtpRenameFile(from, to string, model *FtpHelpStruct) error {
	var (
		err error
	)
	if ftpConntion == nil || ftpConntion.NoOp() != nil {
		ftpConntion, err = ftpLogin(model)
		if err != nil {
			return err
		}
	}
	return ftpConntion.Rename(from, to)
}

// 获取FTP上所有的文件列表
// create by gloomy 2017-04-12 10:43:25
// 目录地址
// 文件列表集 错误对象
func FtpNameList(pathStr string, model *FtpHelpStruct) ([]string, error) {
	var (
		err error
	)
	if ftpConntion == nil || ftpConntion.NoOp() != nil {
		ftpConntion, err = ftpLogin(model)
		if err != nil {
			return err
		}
	}
	return ftpConntion.NameList(pathStr)
}
