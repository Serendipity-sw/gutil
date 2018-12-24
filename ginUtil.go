//gin公共类
//create by gloomy 2017-09-01 01:11:28
package gutil

//gin初始化
//create by gloomy 2017-09-01 01:13:52
//func GinInit(debug bool, serverListeningPort int, rootPrefix string, setGinRouter func(r *gin.Engine, rootPrefix string)) *gin.Engine {
//	gin.SetMode(If(debug, gin.DebugMode, gin.ReleaseMode).(string))
//	rt := gin.Default()
//	setGinRouter(rt, rootPrefix)
//	go rt.Run(fmt.Sprintf(":%d", serverListeningPort))
//	return rt
//}
