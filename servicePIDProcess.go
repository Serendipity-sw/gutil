/**
程序运行pid目录处理类
创建人：邵炜
创建时间：2017年03月11日11:03:55
*/
package gutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const ProgramServicePIDPath = "./programRunPID.pid" // PID文件生成路径

/**
写PID文件
创建人：邵炜
创建时间：2017年03月11日15:47:11
输入参数：文件路径
*/
func WritePid(pidFileStr string) {
	if strings.TrimSpace(pidFileStr) == "" {
		pidFileStr = ProgramServicePIDPath
	}
	pid := os.Getpid()
	f, err := os.OpenFile(pidFileStr, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Open pid file %s failed: %s\n", pidFileStr, err.Error()))
	}

	_, err = f.Write([]byte(fmt.Sprint(pid)))
	if err != nil {
		panic(fmt.Sprintf("Write Pid file %s failed: %s\n", pidFileStr, err.Error()))
	}
	f.Close()
}

/**
检查pid文件是否存在，pid文件中的进程是否存在
创建人：邵炜
创建时间：2017年03月11日15:36:21
输入参数：pid文件路径
输出参数：bool类型（true： 文件不存在或者进程不存在 false: 进程已存在）
*/
func CheckPid(pidFileStr string) bool {
	if strings.TrimSpace(pidFileStr) == "" {
		pidFileStr = ProgramServicePIDPath
	}
	f, err := os.Open(pidFileStr)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(fmt.Sprintf("Open Pid file %s failed: %s\n", pidFileStr, err.Error()))
		}
	}
	defer f.Close()
	// 读取文件内容
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Sprintf("Read Pid file %s failed: %s\n", pidFileStr, err.Error()))
	}
	pid, err := strconv.Atoi(string(buf))
	if err != nil {
		panic(fmt.Sprintf("Convert pid %s failed: %s\n", pid, err.Error()))
	}
	// 进程是否存在
	exist := isProcessExist(pid)
	if exist == false {
		return false
	}
	fmt.Printf("Process with Pid %d is running, exit.\n", pid)
	return true
}

/**
判断进程是否存在
创建人：邵炜
创建时间：2017年03月11日15:31:34
输入参数：进程ID
*/
func isProcessExist(pid int) bool {
	_, err := os.FindProcess(pid)
	return err == nil
}

/**
删除PID文件
创建人：邵炜
创建时间：2017年03月11日15:35:20
输入参数：pid文件路径
*/
func RmPidFile(pidFileStr string) {
	if strings.TrimSpace(pidFileStr) == "" {
		pidFileStr = ProgramServicePIDPath
	}
	err := os.Remove(pidFileStr)
	if err != nil {
		fmt.Printf("rmPidFile remove pidFile is error. err: %s \n", err.Error())
	}
}
