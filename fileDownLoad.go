//文件下载类
//create by gloomy 2017-08-28 13:12:12
package gutil

import (
	"io/ioutil"
	"net/http"
)

//get文件下载
//create by gloomy 2017-08-28 15:33:17
func HttpGetDownFile(urlPathStr, saveFilePath string) error {
	request, err := http.Get(urlPathStr)
	if err != nil {
		return err
	}
	defer request.Body.Close()
	requestByte, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(saveFilePath, requestByte, 0644)
}

//断点续传
//create by gloomy 2017-08-29 16:14:01
func FileTransferProtocol() {

}
