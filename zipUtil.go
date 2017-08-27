//压缩包解压
//create by gloomy 2017-08-27 14:18:42
package gutil

import (
	"os"
	"github.com/c4milo/unpackit"
)

//压缩文件解压
//create by gloomy 2017-08-27 14:21:43
func UnZip(zipPathStr,dirPathStr string) (string,error) {
	file, err := os.Open(zipPathStr)
	if err != nil {
		return "",err
	}
	return unpackit.Unpack(file, dirPathStr)
}
