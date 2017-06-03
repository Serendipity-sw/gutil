// 文件HTTP传输方法
package gutil

import (
	"bytes"
	"github.com/smtc/glog"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

/**
文件发送处理方法
创建人:邵炜
创建时间:2016年11月29日15:37:06
输入参数:文件路径
输出参数:错误对象00001*/
func HttpSendFile(sendHttpUrl, filePathStr, uploadFile string) error {
	fileSizeInfo, err := os.Stat(filePathStr)
	if err != nil {
		return err
	}
	glog.Info("sendFileProcess file size: %d filePathStr: %s \n", fileSizeInfo.Size(), filePathStr)
	extraParams := map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets",
	}
	request, err := newfileUploadRequest(sendHttpUrl, extraParams, uploadFile, filePathStr)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			return err
		}
	}
	return nil
}
