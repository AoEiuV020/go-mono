package main

import (
	"fmt"
	
	"github.com/AoEiuV020/go-mono/packages/common"
	"github.com/AoEiuV020/go-mono/packages/mathlib"
	"github.com/AoEiuV020/go-mono/packages/stringlib"
)

func main() {
	// 创建 logger
	logger := common.NewLogger("DynamicApp")
	logger.Log("应用启动 (动态链接)")
	
	// 使用 mathlib
	calc := mathlib.NewCalculator()
	sum := calc.Add(10, 20)
	product := calc.Multiply(5, 7)
	factorial := calc.Factorial(5)
	maxNum := calc.MaxOfThree(15, 8, 23)
	
	// 使用 stringlib
	processor := stringlib.NewStringProcessor()
	reversed := processor.Reverse("Hello World")
	concatenated := processor.Concat(" - ", "Go", "Mono", "Project")
	upperCase := processor.ToUpperCaseWithLog("golang")
	wordCount := processor.CountWords("This is a test string")
	
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
	
	logger.Log("应用结束")
}
