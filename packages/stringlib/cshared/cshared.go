package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var (
	processors   = make(map[int]*StringProcessor)
	nextProcID   = 1
	processorsMu sync.Mutex
)

//export StringProcessorNew
func StringProcessorNew() C.int {
	processorsMu.Lock()
	defer processorsMu.Unlock()
	
	sp := &StringProcessor{prefix: "StringLib"}
	id := nextProcID
	processors[id] = sp
	nextProcID++
	return C.int(id)
}

//export StringProcessorReverse
func StringProcessorReverse(spID C.int, s *C.char) *C.char {
	processorsMu.Lock()
	sp := processors[int(spID)]
	processorsMu.Unlock()
	
	if sp != nil {
		result := sp.Reverse(C.GoString(s))
		return C.CString(result)
	}
	return nil
}

//export StringProcessorConcat
func StringProcessorConcat(spID C.int, sep *C.char, parts **C.char, count C.int) *C.char {
	processorsMu.Lock()
	sp := processors[int(spID)]
	processorsMu.Unlock()
	
	if sp != nil {
		// 将 C 字符串数组转换为 Go 字符串切片
		goSlice := (*[1 << 28]*C.char)(unsafe.Pointer(parts))[:count:count]
		goParts := make([]string, count)
		for i := 0; i < int(count); i++ {
			goParts[i] = C.GoString(goSlice[i])
		}
		
		result := sp.Concat(C.GoString(sep), goParts...)
		return C.CString(result)
	}
	return nil
}

//export StringProcessorToUpperCase
func StringProcessorToUpperCase(spID C.int, s *C.char) *C.char {
	processorsMu.Lock()
	sp := processors[int(spID)]
	processorsMu.Unlock()
	
	if sp != nil {
		result := sp.ToUpperCaseWithLog(C.GoString(s))
		return C.CString(result)
	}
	return nil
}

//export StringProcessorCountWords
func StringProcessorCountWords(spID C.int, s *C.char) C.int {
	processorsMu.Lock()
	sp := processors[int(spID)]
	processorsMu.Unlock()
	
	if sp != nil {
		return C.int(sp.CountWords(C.GoString(s)))
	}
	return 0
}

//export FreeString
func FreeString(s *C.char) {
	C.free(unsafe.Pointer(s))
}

// StringProcessor 提供字符串处理功能
type StringProcessor struct {
	prefix string
}

// Reverse 反转字符串
func (sp *StringProcessor) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	result := string(runes)
	sp.log(fmt.Sprintf("Reverse(%q) = %q", s, result))
	return result
}

// Concat 连接多个字符串
func (sp *StringProcessor) Concat(separator string, parts ...string) string {
	result := strings.Join(parts, separator)
	sp.log(fmt.Sprintf("Concat with separator %q: %d parts", separator, len(parts)))
	return result
}

// ToUpperCaseWithLog 将字符串转换为大写
func (sp *StringProcessor) ToUpperCaseWithLog(s string) string {
	result := toUpperCase(s)
	sp.log(fmt.Sprintf("ToUpperCase(%q) = %q", s, result))
	return result
}

// CountWords 计算字符串中的单词数
func (sp *StringProcessor) CountWords(s string) int {
	words := strings.Fields(s)
	count := len(words)
	sp.log(fmt.Sprintf("CountWords(%q) = %d", s, count))
	return count
}

func (sp *StringProcessor) log(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", timestamp, sp.prefix, message)
}

func toUpperCase(s string) string {
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

func main() {}
