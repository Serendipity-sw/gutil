// 压缩包解压
// create by gloomy 2017-08-27 14:18:42
package gutil

import (
	"compress/gzip"
	"io/ioutil"
	"os"
)

//压缩文件解压 (方法暂时遗弃,解压会出现数据丢失)
//create by gloomy 2017-08-27 14:21:43
//func UnZip(zipPathStr, dirPathStr string) (string, error) {
//	file, err := os.Open(zipPathStr)
//	if err != nil {
//		return "", err
//	}
//	return unpackit.Unpack(file, dirPathStr)
//}

// gz文件解压
// create by gloomy 2017-09-04 20:27:24
func UnGzip(zipPathStr, filePathStr string) error {
	f, err := os.Open(zipPathStr)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer w.Close()
	readContentByte, err := ioutil.ReadAll(w)
	if err != nil {
		return err
	}
	return FileCreateAndWrite(&readContentByte, filePathStr, false)
}
