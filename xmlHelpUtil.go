package gutil

import "strings"

var ReplaceXmlStrArray map[string]string = make(map[string]string)

func init() {

	ReplaceXmlStrArray["&"] = "&amp;"
	ReplaceXmlStrArray["\""] = "&quot;"
	ReplaceXmlStrArray["<"] = "&lt"
	ReplaceXmlStrArray[">"] = "&gt"
	ReplaceXmlStrArray["'"] = "&apos;"
}

// 生成xml文件修正xml节点内容
// create by gloomysw 2017-5-25 16:18:11
func XmlContentReplace(value string) string {
	for key, replaceStr := range ReplaceXmlStrArray {
		value = strings.Replace(value, key, replaceStr, -1)
	}
	return value
}
