// 文件监控类
// create by gloomy 2017-5-3 11:40:32
package gutil

import (
	"errors"
	"fmt"
	"github.com/gobwas/glob"
	"github.com/howeyc/fsnotify"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	AutoMatedTaskLock sync.RWMutex
	AutoMatedTaskFile map[string]int = make(map[string]int)
)

// 文件监控
// create by gloomy 2017-5-3 11:42:09
func WatchFile(filePathStr, matchFileName string, deleteFileCallBack, modifyFileCallBack, renameFileCallBack, createFileCallBack func(string)) (*fsnotify.Watcher, error) {
	fi, err := os.Stat(filePathStr)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, errors.New("WatchFile filePathStr isn't dir!")
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev == nil {
					continue
				}
				if strings.TrimSpace(matchFileName) != "" {
					if glob.MustCompile(matchFileName).Match(ev.Name) {
						fmt.Printf("WatchFile file is not match! matchFileName: %s fileName: %s \n", matchFileName, ev.Name)
						continue
					}
				}
				if ev.IsDelete() && deleteFileCallBack != nil {
					deleteFileCallBack(ev.Name)
				}
				if ev.IsRename() && renameFileCallBack != nil {
					renameFileCallBack(ev.Name)
				}
				if ev.IsCreate() && createFileCallBack != nil {
					createFileCallBack(ev.Name)
				}
				if ev.IsModify() && modifyFileCallBack != nil {
					AutoMatedTaskLock.RLock()
					_, ok := AutoMatedTaskFile[ev.Name]
					AutoMatedTaskLock.RUnlock()
					if !ok {
						AutoMatedTaskLock.Lock()
						AutoMatedTaskFile[ev.Name] = 0
						AutoMatedTaskLock.Unlock()
						go WatchFileAutoMated(ev.Name, modifyFileCallBack)
					}
				}
			case err := <-watcher.Error:
				if err == nil {
					continue
				}
				fmt.Printf("WatchFile fsnotify watcher is error! err: %s \n", err.Error())
			}
		}
	}()
	err = watcher.WatchFlags(filePathStr, fsnotify.FSN_ALL)
	if err != nil {
		return nil, err
	}
	return watcher, err
}

// 自动化创建任务 需要监控的文件,判断文件是否上传完毕
// 创建人:邵炜
// 创建时间:2016年9月5日14:58:13
// 输入参数: 文件路劲
func WatchFileAutoMated(filePath string, callBack func(string)) {
	tmrIntal := 30 * time.Second
	fileSaveTmr := time.NewTimer(tmrIntal)
	fileState, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("watchFileAutoMated can't load file! path: %s err: %s \n", filePath, err.Error())
		return
	}
	var (
		size   = fileState.Size()
		number int64
	)
	defer func() {
		fileSaveTmr.Stop()
		AutoMatedTaskLock.Lock()
		delete(AutoMatedTaskFile, filePath)
		AutoMatedTaskLock.Unlock()
	}()
	<-fileSaveTmr.C
	for {
		fileState, err = os.Stat(filePath)
		if err != nil {
			fmt.Printf("watchFileAutoMated can't load file! path: %s err: %s \n", filePath, err.Error())
			return
		}
		number = fileState.Size()
		if size == number {
			go callBack(filePath)
			return
		}
		size = number
		fileSaveTmr.Reset(tmrIntal)
		<-fileSaveTmr.C
	}
}
