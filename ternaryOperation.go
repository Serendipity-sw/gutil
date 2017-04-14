/**
公共组件库
构建三元运算
创建人:邵炜
创建时间:2017年2月8日18:50:51
*/
package gutil

/**
构建三元运算
创建人:邵炜
创建时间:2017年2月8日18:51:36
输入参数:是否匹配 第一返回数 第二返回数
*/
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
