package stringlib

import (
	"strings"
	
	"github.com/AoEiuV020/go-mono/packages/common"
)

// StringProcessor 提供字符串处理功能
type StringProcessor struct {
	logger *common.Logger
}

// NewStringProcessor 创建一个新的 StringProcessor 实例
func NewStringProcessor() *StringProcessor {
	return &StringProcessor{
		logger: common.NewLogger("StringLib"),
	}
}

// Reverse 反转字符串
func (sp *StringProcessor) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	result := string(runes)
	sp.logger.LogFormat("Reverse(%q) = %q", s, result)
	return result
}

// Concat 连接多个字符串
func (sp *StringProcessor) Concat(separator string, parts ...string) string {
	result := strings.Join(parts, separator)
	sp.logger.LogFormat("Concat with separator %q: %d parts", separator, len(parts))
	return result
}

// ToUpperCaseWithLog 将字符串转换为大写，使用 common 包的函数
func (sp *StringProcessor) ToUpperCaseWithLog(s string) string {
	result := common.ToUpperCase(s)
	sp.logger.LogFormat("ToUpperCase(%q) = %q", s, result)
	return result
}

// CountWords 计算字符串中的单词数
func (sp *StringProcessor) CountWords(s string) int {
	words := strings.Fields(s)
	count := len(words)
	sp.logger.LogFormat("CountWords(%q) = %d", s, count)
	return count
}
