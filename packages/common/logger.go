package common

import (
	"fmt"
	"time"
)

// Logger 提供简单的日志功能
type Logger struct {
	prefix string
}

// NewLogger 创建一个新的 Logger 实例
func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

// Log 输出日志信息
func (l *Logger) Log(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", timestamp, l.prefix, message)
}

// LogFormat 输出格式化的日志信息
func (l *Logger) LogFormat(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.Log(message)
}
