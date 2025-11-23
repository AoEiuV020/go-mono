package common

import "strings"

// ToUpperCase 将字符串转换为大写
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// Max 返回两个整数中的较大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min 返回两个整数中的较小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
