/**
随机数处理
create by gloomy 2017-03-29 22:11:23
*/
package common

import "math"

/**
四舍五入取舍
create by gloomy 2017-03-29 22:11:18
需要取舍的浮点数 取舍几位
*/
func Rounding(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
