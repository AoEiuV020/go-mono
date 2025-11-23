package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"sync"
	"unsafe"
	
	"github.com/AoEiuV020/go-mono/packages/common"
)

var (
	loggers   = make(map[int]*common.Logger)
	nextID    = 1
	loggersMu sync.Mutex
)

//export LoggerNew
func LoggerNew(prefix *C.char) C.int {
	loggersMu.Lock()
	defer loggersMu.Unlock()
	
	// 调用原始 Go 代码
	logger := common.NewLogger(C.GoString(prefix))
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
		// 调用原始 Go 代码
		logger.Log(C.GoString(message))
	}
}

//export LoggerLogFormat
func LoggerLogFormat(loggerID C.int, format *C.char, args *C.char) {
	loggersMu.Lock()
	logger := loggers[int(loggerID)]
	loggersMu.Unlock()
	
	if logger != nil {
		// 简单处理：直接输出格式化后的字符串
		logger.Log(C.GoString(args))
	}
}

//export ToUpperCase
func ToUpperCase(s *C.char) *C.char {
	// 调用原始 Go 代码
	result := common.ToUpperCase(C.GoString(s))
	return C.CString(result)
}

//export Max
func Max(a, b C.int) C.int {
	// 调用原始 Go 代码
	return C.int(common.Max(int(a), int(b)))
}

//export Min
func Min(a, b C.int) C.int {
	// 调用原始 Go 代码
	return C.int(common.Min(int(a), int(b)))
}

//export FreeString
func FreeString(s *C.char) {
	C.free(unsafe.Pointer(s))
}

func main() {}
