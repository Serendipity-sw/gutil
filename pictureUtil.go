package gutil

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

/**
base64图片转图片文件
创建人:邵炜
创建时间:2016年12月26日17:12:22
输入参数:图片base内容 图片文件存放路径(不包含图片名称)
输出参数:图片名称 错误对象
*/
func PictureBase64ToFile(fileContentTem *string) (string, error) {
	fileContent := *fileContentTem
	fileContentArray := strings.Split(fileContent, ",")
	if len(fileContentArray) != 2 {
		return "", errors.New("picture base64 content error!")
	}
	dataTypeArray := strings.Split(fileContentArray[0], ";")
	if len(dataTypeArray) != 2 {
		return "", errors.New("picture dataType error!")
	}
	fileTypeArray := strings.Split(dataTypeArray[0], "/")
	if len(fileTypeArray) != 2 {
		return "", errors.New("picture type name error!")
	}
	fileType := fileTypeArray[1]
	if strings.ToLower(fileType) == "jpeg" {
		fileType = "jpg"
	}
	dist, _ := base64.StdEncoding.DecodeString(fileContentArray[1])
	fileName := fmt.Sprintf("./%d.%s", time.Now().UnixNano(), fileType)
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	f.Write(dist)
	f.Close()
	return fileName, nil
}
