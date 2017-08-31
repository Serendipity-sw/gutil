//数组操作类
//create by gloomy 2017-08-31 00:10:08
package gutil

//字符串数组去重
//create by gloomy 2017-08-31 00:10:32
func UniqueSlice(slice *[]string) {
	found := make(map[string]bool)
	total := 0
	for i, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[i]
			total++
		}
	}
	*slice = (*slice)[:total]
}
