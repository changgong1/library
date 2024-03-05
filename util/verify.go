package util

import "strings"

func IsAlphaNumeric(str string) bool {
	// 遍历字符串，判断每个字符是否为字母数字
	for _, ch := range str {
		if !strings.ContainsRune("0123456789", ch) {
			return false
		}
	}
	return true
}
