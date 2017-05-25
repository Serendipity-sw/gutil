package gutil

import "strings"

var replaceXmlStrArray map[string]string = make(map[string]string)

func init() {
	replaceXmlStrArray = append("&", "&amp;")
	replaceXmlStrArray = append("\"", "&quot;")
	replaceXmlStrArray = append("<", "&lt")
	replaceXmlStrArray = append(">", "&gt")
	replaceXmlStrArray = append("'", "&apos;")
}

// 生成xml文件修正xml节点内容
// create by gloomysw 2017-5-25 16:18:11
func XmlContentReplace(value string) string {
	for key, replaceStr := range replaceXmlStrArray {
		value = strings.Replace(value, key, replaceStr, -1)
	}
	return value
}
