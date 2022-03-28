// Package word 实现单词类型转换功能：
// - 单词全部转换为小写
// - 单词全部转换为大写
// - 下划线单词转换为大写驼峰单词
// - 下划线单词转换为小写驼峰单词
// - 驼峰单词转为下划线单词
package word

import (
	"strings"
	"unicode"
)

// ToUpper 单词全部转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 单词全部转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCameCase 下划线单词转换为大写驼峰单词
func UnderscoreToUpperCameCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线单词转换为小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCameCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰单词转为下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
