package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"sync"
	"time"
)

var (
	calculators   = make(map[int]*Calculator)
	nextCalcID    = 1
	calculatorsMu sync.Mutex
)

//export CalculatorNew
func CalculatorNew() C.int {
	calculatorsMu.Lock()
	defer calculatorsMu.Unlock()
	
	calc := &Calculator{prefix: "MathLib"}
	id := nextCalcID
	calculators[id] = calc
	nextCalcID++
	return C.int(id)
}

//export CalculatorAdd
func CalculatorAdd(calcID C.int, a, b C.int) C.int {
	calculatorsMu.Lock()
	calc := calculators[int(calcID)]
	calculatorsMu.Unlock()
	
	if calc != nil {
		return C.int(calc.Add(int(a), int(b)))
	}
	return 0
}

//export CalculatorMultiply
func CalculatorMultiply(calcID C.int, a, b C.int) C.int {
	calculatorsMu.Lock()
	calc := calculators[int(calcID)]
	calculatorsMu.Unlock()
	
	if calc != nil {
		return C.int(calc.Multiply(int(a), int(b)))
	}
	return 0
}

//export CalculatorFactorial
func CalculatorFactorial(calcID C.int, n C.int) C.int {
	calculatorsMu.Lock()
	calc := calculators[int(calcID)]
	calculatorsMu.Unlock()
	
	if calc != nil {
		return C.int(calc.Factorial(int(n)))
	}
	return 0
}

//export CalculatorMaxOfThree
func CalculatorMaxOfThree(calcID C.int, a, b, c C.int) C.int {
	calculatorsMu.Lock()
	calc := calculators[int(calcID)]
	calculatorsMu.Unlock()
	
	if calc != nil {
		return C.int(calc.MaxOfThree(int(a), int(b), int(c)))
	}
	return 0
}

// Calculator 提供数学计算功能
type Calculator struct {
	prefix string
}

// Add 执行加法运算
func (c *Calculator) Add(a, b int) int {
	result := a + b
	c.log(fmt.Sprintf("Add(%d, %d) = %d", a, b, result))
	return result
}

// Multiply 执行乘法运算
func (c *Calculator) Multiply(a, b int) int {
	result := a * b
	c.log(fmt.Sprintf("Multiply(%d, %d) = %d", a, b, result))
	return result
}

// Factorial 计算阶乘
func (c *Calculator) Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	c.log(fmt.Sprintf("Factorial(%d) = %d", n, result))
	return result
}

// MaxOfThree 返回三个数中的最大值
func (c *Calculator) MaxOfThree(a, b, num int) int {
	result := max(max(a, b), num)
	c.log(fmt.Sprintf("MaxOfThree(%d, %d, %d) = %d", a, b, num, result))
	return result
}

func (c *Calculator) log(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", timestamp, c.prefix, message)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {}
