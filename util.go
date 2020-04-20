package gexport

import (
	"bytes"
	"strings"
)

//NameCamelCase 转换为驼峰命名
func NameCamelCase(str string) string {
	var buf bytes.Buffer
	strSlices := strings.Split(str, "_")
	if len(strSlices) == 0 {
		return Capitalize(str)
	}
	for _, s := range strSlices {
		buf.WriteString(Capitalize(s))
	}
	return buf.String()
}

//Capitalize 首字符大写
func Capitalize(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
