// 文件数据记录类
// create by gloomy 2017-04-06 10:11:35
package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const maxFileDataRecordingBytes = 1000000 // 默认文件大小

// 文件数据记录对象
// create by gloomy 2017-04-06 10:15:00
type FileDataRecording struct {
	sync.Mutex                         // 锁
	F                         *os.File // 文件对象
	FilePre                   string   // 文件开头字符串
	Fn                        string   // 文件路径
	Bytes                     int      // 文件大小
	Seq                       int      // 第几个
	FileProgram               string   // 文件存放路径
	MaxFileDataRecordingBytes int      // 文件大小
}

// 打开文件数据记录
// create by gloomy 2017-04-06 10:17:38
// 文件存放目录地址 文件开头字符串 文件大小
// 文件数据对象
func OpenLoadFile(fileProgram, filePre string, maxSize int) *FileDataRecording {
	if maxSize == 0 {
		maxSize = maxFileDataRecordingBytes
	}
	lf := &FileDataRecording{
		FilePre:                   filePre,
		FileProgram:               fileProgram,
		MaxFileDataRecordingBytes: maxSize,
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

// 创建新文件
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
	if f.Bytes+dataStrLen > f.MaxFileDataRecordingBytes {
		f.Close()
		if err = f.CreateNewFile(); err != nil {
			return
		}
	}
	f.Bytes += dataStrLen
	_, err = f.F.WriteString(dataStr)
	return
}

// 获取所有完成的文件列表
// create by gloomy 2017-04-06 13:46:51
// 文件列表
func (f *FileDataRecording) FileList() *[]string {
	var (
		fileArray []string
	)
	filepath.Walk(f.FileProgram, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if ext := filepath.Ext(path); ext == ".tmp" {
			return nil
		}
		if info.Size() == 0 {
			os.Remove(path)
			return nil
		}
		if strings.HasPrefix(path, f.FilePre) {
			fileArray = append(fileArray, path)
		}
		return nil
	})
	return &fileArray
}

// 删除过期文件
// create by gloomy 2017-04-06 22:53:17
// 几天前
func (f *FileDataRecording) FileListRemoveOld(days int) {
	var (
		timeDate   = time.Now().AddDate(0, 0, 0-days).UnixNano()
		unixNumber int
		fileName   []string
	)
	filepath.Walk(f.FileProgram, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if ext := filepath.Ext(path); ext == ".tmp" {
			return nil
		}
		if info.Size() == 0 {
			os.Remove(path)
			return nil
		}
		if strings.HasPrefix(path, f.FilePre) {
			fileName = strings.Split(path, "-")
			if len(fileName) == 3 {
				unixNumber, err = strconv.Atoi(fileName[1])
				if err != nil {
					fmt.Printf("FileListRemoveOld fileProgram: %s path: %s err: %s \n", f.FileProgram, path, err.Error())
					return nil
				}
				if timeDate >= int64(unixNumber) {
					os.Remove(path)
				}
			}
		}
		return nil
	})
}
