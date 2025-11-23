package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"sync"
	
	"github.com/AoEiuV020/go-mono/packages/mathlib"
)

var (
	calculators   = make(map[int]*mathlib.Calculator)
	nextCalcID    = 1
	calculatorsMu sync.Mutex
)

//export CalculatorNew
func CalculatorNew() C.int {
	calculatorsMu.Lock()
	defer calculatorsMu.Unlock()
	
	// 调用原始 Go 代码
	calc := mathlib.NewCalculator()
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
		// 调用原始 Go 代码
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
		// 调用原始 Go 代码
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
		// 调用原始 Go 代码
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
		// 调用原始 Go 代码
		return C.int(calc.MaxOfThree(int(a), int(b), int(c)))
	}
	return 0
}

func main() {}
