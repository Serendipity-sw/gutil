/**
时间处理类
创建人：邵炜
创建时间：2017年03月21日20:53:46
*/
package common

import (
	"strings"
	"time"
)

/**
时间格式化处理
创建人:邵炜
创建时间:2017年3月22日09:13:01
输入参数:需要格式化的时间 格式化方式 示例yyyy-MM-dd hh:mm:ss.tttttt   2017-03-22 10:21:55.379415
*/
func DateFormat(timeDate time.Time, layerout string) string {
	var (
		mapList     map[string]string = make(map[string]string)
		millisecond                   = strings.Replace(timeDate.Format(".999999"), ".", "", -1)
	)
	mapList["yyyy"] = timeDate.Format("2006")
	mapList["MM"] = timeDate.Format("01")
	mapList["dd"] = timeDate.Format("02")
	mapList["hh"] = timeDate.Format("15")
	mapList["mm"] = timeDate.Format("04")
	mapList["ss"] = timeDate.Format("05")
	mapList["tttttt"] = millisecond
	if len(millisecond) >= 5 {
		mapList["ttttt"] = millisecond[:5]
	}
	if len(millisecond) >= 4 {
		mapList["tttt"] = millisecond[:4]
	}
	if len(millisecond) >= 3 {
		mapList["ttt"] = millisecond[:3]
	}
	if len(millisecond) >= 2 {
		mapList["tt"] = millisecond[:2]
	}
	if len(millisecond) >= 1 {
		mapList["t"] = millisecond[:1]
	}
	for key, value := range mapList {
		layerout = strings.Replace(layerout, key, value, -1)
	}
	return layerout
}
