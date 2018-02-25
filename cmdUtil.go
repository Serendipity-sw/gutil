package gutil

import (
	"bufio"
	"io"
	"os/exec"
)

//执行函数命令 commandContent 需要执行的命令
//create by gloomy 2018-2-25 16:41:52
func ExecCommand(commandContent string) (*[]string, error) {
	var resultContentArray []string
	cmd := exec.Command("cmd", "/C", commandContent)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return &resultContentArray, err
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		resultContentArray = append(resultContentArray, line)
	}
	cmd.Wait()
	return &resultContentArray, nil
}
