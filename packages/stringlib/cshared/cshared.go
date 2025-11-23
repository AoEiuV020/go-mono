package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"sync"
	"unsafe"
	
	"github.com/AoEiuV020/go-mono/packages/stringlib"
)

var (
	processors   = make(map[int]*stringlib.StringProcessor)
	nextProcID   = 1
	processorsMu sync.Mutex
)

//export StringProcessorNew
func StringProcessorNew() C.int {
	processorsMu.Lock()
	defer processorsMu.Unlock()
	
	// 调用原始 Go 代码
	sp := stringlib.NewStringProcessor()
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
		// 调用原始 Go 代码
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
		
		// 调用原始 Go 代码
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
		// 调用原始 Go 代码
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
		// 调用原始 Go 代码
		return C.int(sp.CountWords(C.GoString(s)))
	}
	return 0
}

//export FreeString
func FreeString(s *C.char) {
	C.free(unsafe.Pointer(s))
}

func main() {}
