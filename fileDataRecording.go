// 文件数据记录类
// create by gloomy 2017-04-06 10:11:35
package common

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const maxFileDataRecordingBytes = 1000000

// 文件数据记录对象
// create by gloomy 2017-04-06 10:15:00
type FileDataRecording struct {
	sync.Mutex           // 锁
	F           *os.File // 文件对象
	FilePre     string   // 文件开头字符串
	Fn          string   // 文件路径
	Bytes       int      // 文件大小
	Seq         int      // 第几个
	FileProgram string   // 文件存放路径
}

// 打开文件数据记录
// create by gloomy 2017-04-06 10:17:38
// 文件存放目录地址 文件开头字符串
// 文件数据对象
func OpenLoadFile(fileProgram, filePre string) *FileDataRecording {
	lf := &FileDataRecording{
		FilePre:     filePre,
		FileProgram: fileProgram,
	}
	lf.Rotate()
	return lf
}

// 文件退出
// create by gloomy 2017-04-06 10:27:58
func (f *FileDataRecording) Exit() {
	f.Lock()
	f.Close()
	f.F = nil
	f.Unlock()
}

// 文件关闭
// create by gloomy 2017-04-06 10:22:14
func (f *FileDataRecording) Close() {
	if f.F != nil {
		f.F.Close()
		os.Rename(f.Fn, f.Fn[0:len(f.Fn)-4]) //去掉末尾的.tmp
	}
}

// 文件切换
// create by gloomy 2017-04-06 10:30:05
func (f *FileDataRecording) Rotate() {
	f.Lock()
	f.Seq = 0
	f.Close()
	f.CreateNewFile()
	f.Unlock()
}

/// 创建新文件
// create by gloomy 2017-04-06 10:33:11
// 错误对象
func (f *FileDataRecording) CreateNewFile() (err error) {
	f.Bytes = 0
	if !strings.HasSuffix(f.FileProgram, "/") {
		f.FileProgram += "/"
	}
	f.Fn = fmt.Sprintf("%s%s-%d-%d.tmp", f.FileProgram, f.FilePre, time.Now().UnixNano(), f.Seq)
	f.Seq++
	f.F, err = os.OpenFile(f.Fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create file %s failed: %s \n", f.Fn, err.Error())
	}
	return
}

// 写入数据
// create by gloomy 2017-04-06 11:40:55
// 需要写入的数据
// 错误对象
func (f *FileDataRecording) WriteData(dataStr string) (err error) {
	f.Lock()
	defer f.Unlock()
	if f.F == nil {
		err = f.CreateNewFile()
		if err != nil {
			return
		}
	}
	dataStrLen := len(dataStr)
	if f.Bytes+dataStrLen > maxFileDataRecordingBytes {
		f.Close()
		if err = f.CreateNewFile(); err != nil {
			return
		}
	}
	f.Bytes += dataStrLen
	_, err = f.F.WriteString(dataStr)
	return
}
