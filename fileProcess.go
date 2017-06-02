/**
公共组件库
文件处理
创建人:邵炜
创建时间:2017年2月8日18:11:29
*/
package gutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/**
根据文件夹路径创建文件,如文件存在则不做任何操作
创建人:邵炜
创建时间:2016年12月21日17:23:54
输入参数:文件夹路径
输出参数:错误对象
*/
func CreateFileProcess(path string) error {
	fileExists, err := PathExists(path)
	if err != nil {
		return err
	}
	if !fileExists {
		err = os.MkdirAll(path, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

/**
判断文件或文件夹是否存在
创建人:邵炜
创建时间:2016年12月21日17:07:42
输入参数:需要查询的文件或文件夹路径
输出参数:返回值true存在 否则不存在  错误对象
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return false, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
写文件
创建人:邵炜
创建时间:2016年9月7日16:31:39
输入参数:文件内容 写入文件的路劲(包含文件名) 是否追加写入
输出参数:错误对象
*/
func FileCreateAndWrite(content *[]byte, fileName string, isAppend bool) error {
	var (
		f   *os.File
		err error
	)
	if isAppend {
		f, err = AppendFileOpen(fileName)
	} else {
		f, err = FileOpen(fileName)
	}
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(*content)
	if err != nil {
		return err
	}
	return nil
}

/**
文件读取逐行进行读取
创建人:邵炜
创建时间:2016年9月20日10:23:41
输入参数: 文件路劲
输出参数: 字符串数组(数组每一项对应文件的每一行) 错误对象
*/
func ReadFileByLine(filePath string) (*[]string, error) {
	var (
		readAll     = false
		readByte    []byte
		line        []byte
		err         error
		contentLine []string
	)
	fs, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs)
	for err != io.EOF {
		if err != nil {
		}
		if readAll {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
		} else {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
			if len(strings.TrimSpace(string(line))) == 0 {
				continue
			}
			contentLine = append(contentLine, string(line))
			line = line[:0]
		}
	}
	return &contentLine, nil
}

// 读取文件行数
// create by gloomy 2017-6-2 14:00:32
func ReadFileLineNumber(filePathStr string) (int, error) {
	var (
		readAll           = false
		readByte          []byte
		line              []byte
		err               error
		contentLineNumber = 0
	)
	fs, err := os.Open(filePathStr)
	if err != nil {
		return contentLineNumber, err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs)
	for err != io.EOF {
		if err != nil {
		}
		if readAll {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
		} else {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
			if len(strings.TrimSpace(string(line))) == 0 {
				continue
			}
			contentLineNumber++
			line = line[:0]
		}
	}
	return contentLineNumber, nil
}

/**
根据条件读文件
创建人:邵炜
创建时间:2017年3月22日11:03:31
输入参数:文件路径 文件写入对象 条件平判断方法
输出参数:错误对象
*/
func RWFileByWhere(fileName string, fileWrite *os.File, where func(content string, fileWrite *os.File)) error {
	var (
		readAll  = false
		readByte []byte
		line     []byte
		err      error
	)
	fs, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs)
	for err != io.EOF {
		if err != nil {
		}
		if readAll {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
		} else {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
			if len(strings.TrimSpace(string(line))) == 0 {
				continue
			}
			where(string(line), fileWrite)
			line = line[:0]
		}
	}
	return nil
}

/**
文件打开
创建人:邵炜
创建时间:2017年3月14日14:54:08
输入参数:文件路径
输出参数:文件对象 错误对象
*/
func AppendFileOpen(fileName string) (*os.File, error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return f, err
}

/**
文件打开
创建人:邵炜
创建时间:2017年3月14日14:54:08
输入参数:文件路径
输出参数:文件对象 错误对象
*/
func FileOpen(fileName string) (*os.File, error) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	return f, err
}
