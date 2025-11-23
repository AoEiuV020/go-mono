package main

/*
#cgo LDFLAGS: -L${SRCDIR}/../../build/dynamic/lib -lcommon -lmathlib -lstringlib
#include <stdlib.h>

// common.h
int LoggerNew(char* prefix);
void LoggerLog(int logger, char* message);
char* ToUpperCase(char* s);
int Max(int a, int b);
void FreeString(char* s);

// mathlib.h
int CalculatorNew();
int CalculatorAdd(int calc, int a, int b);
int CalculatorMultiply(int calc, int a, int b);
int CalculatorFactorial(int calc, int n);
int CalculatorMaxOfThree(int calc, int a, int b, int c);

// stringlib.h
int StringProcessorNew();
char* StringProcessorReverse(int sp, char* s);
char* StringProcessorConcat(int sp, char* sep, char** parts, int count);
char* StringProcessorToUpperCase(int sp, char* s);
int StringProcessorCountWords(int sp, char* s);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// 创建 logger
	cPrefix := C.CString("DynamicApp")
	defer C.free(unsafe.Pointer(cPrefix))
	logger := C.LoggerNew(cPrefix)
	
	cMsg := C.CString("应用启动 (动态链接)")
	C.LoggerLog(logger, cMsg)
	C.free(unsafe.Pointer(cMsg))
	
	// 使用 mathlib
	calc := C.CalculatorNew()
	sum := int(C.CalculatorAdd(calc, 10, 20))
	product := int(C.CalculatorMultiply(calc, 5, 7))
	factorial := int(C.CalculatorFactorial(calc, 5))
	maxNum := int(C.CalculatorMaxOfThree(calc, 15, 8, 23))
	
	// 使用 stringlib
	processor := C.StringProcessorNew()
	
	cStr := C.CString("Hello World")
	cReversed := C.StringProcessorReverse(processor, cStr)
	reversed := C.GoString(cReversed)
	C.FreeString(cReversed)
	C.free(unsafe.Pointer(cStr))
	
	// 连接字符串
	parts := []string{"Go", "Mono", "Project"}
	cParts := make([]*C.char, len(parts))
	for i, part := range parts {
		cParts[i] = C.CString(part)
	}
	cSep := C.CString(" - ")
	cConcatenated := C.StringProcessorConcat(processor, cSep, &cParts[0], C.int(len(parts)))
	concatenated := C.GoString(cConcatenated)
	C.FreeString(cConcatenated)
	C.free(unsafe.Pointer(cSep))
	for _, cPart := range cParts {
		C.free(unsafe.Pointer(cPart))
	}
	
	cGoLang := C.CString("golang")
	cUpperCase := C.StringProcessorToUpperCase(processor, cGoLang)
	upperCase := C.GoString(cUpperCase)
	C.FreeString(cUpperCase)
	C.free(unsafe.Pointer(cGoLang))
	
	cTestStr := C.CString("This is a test string")
	wordCount := int(C.StringProcessorCountWords(processor, cTestStr))
	C.free(unsafe.Pointer(cTestStr))
	
	// 输出结果摘要
	fmt.Println("\n========== 计算结果摘要 ==========")
	fmt.Printf("加法结果: %d\n", sum)
	fmt.Printf("乘法结果: %d\n", product)
	fmt.Printf("阶乘结果: %d\n", factorial)
	fmt.Printf("最大值: %d\n", maxNum)
	fmt.Printf("反转字符串: %s\n", reversed)
	fmt.Printf("连接字符串: %s\n", concatenated)
	fmt.Printf("大写字符串: %s\n", upperCase)
	fmt.Printf("单词数: %d\n", wordCount)
	fmt.Println("==================================")
	
	cEndMsg := C.CString("应用结束")
	C.LoggerLog(logger, cEndMsg)
	C.free(unsafe.Pointer(cEndMsg))
}
