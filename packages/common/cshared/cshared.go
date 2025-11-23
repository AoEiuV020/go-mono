package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

var (
	loggers   = make(map[int]*Logger)
	nextID    = 1
	loggersMu sync.Mutex
)

//export LoggerNew
func LoggerNew(prefix *C.char) C.int {
	loggersMu.Lock()
	defer loggersMu.Unlock()
	
	logger := &Logger{prefix: C.GoString(prefix)}
	id := nextID
	loggers[id] = logger
	nextID++
	return C.int(id)
}

//export LoggerLog
func LoggerLog(loggerID C.int, message *C.char) {
	loggersMu.Lock()
	logger := loggers[int(loggerID)]
	loggersMu.Unlock()
	
	if logger != nil {
		logger.Log(C.GoString(message))
	}
}

//export ToUpperCase
func ToUpperCase(s *C.char) *C.char {
	result := toUpperCase(C.GoString(s))
	return C.CString(result)
}

//export Max
func Max(a, b C.int) C.int {
	return C.int(max(int(a), int(b)))
}

//export FreeString
func FreeString(s *C.char) {
	C.free(unsafe.Pointer(s))
}

// Logger 提供简单的日志功能
type Logger struct {
	prefix string
}

// Log 输出日志信息
func (l *Logger) Log(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", timestamp, l.prefix, message)
}

func toUpperCase(s string) string {
	// 简单实现
	result := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			result += string(c - 32)
		} else {
			result += string(c)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {}
