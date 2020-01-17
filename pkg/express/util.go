package express

import "strings"

const (
	empty = ""
	slot  = " "
)

// 清理多余的空格
func cleanDuplicatedSpaces(address string) string {
	return duplicatedSpaceRegexp.ReplaceAllLiteralString(address, slot)
}

// 从字符串删除指定的子字符串
func deleteSubString(src string, sub string) string {
	return strings.ReplaceAll(src, sub, empty)
}

// 统一字符串中的手机号码格式
func unifyPhone(content string) string {
	for _, exp := range phoneCompatibleRegexps {
		content = exp.ReplaceAllLiteralString(content, "$1$2$3")
	}
	return content
}
