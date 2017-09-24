package gutil

import "github.com/smtc/glog"

//日志初始化
//当程序停止前需将glog进行关闭 调用glog.Close方法
//create by gloomy 2017-9-24 14:23:45
func LogInit(debug bool, logsDir string) {
	option := map[string]interface{}{
		"typ": "file",
	}
	if len(logsDir) != 0 {
		option["dir"] = logsDir
	}
	if debug {
		glog.InitLogger(glog.DEV, option)
	} else {
		glog.InitLogger(glog.PRO, option)
	}
}
