package mathlib

import (
	"github.com/AoEiuV020/go-mono/packages/common"
)

// Calculator 提供数学计算功能
type Calculator struct {
	logger *common.Logger
}

// NewCalculator 创建一个新的 Calculator 实例
func NewCalculator() *Calculator {
	return &Calculator{
		logger: common.NewLogger("MathLib"),
	}
}

// Add 执行加法运算
func (c *Calculator) Add(a, b int) int {
	result := a + b
	c.logger.LogFormat("Add(%d, %d) = %d", a, b, result)
	return result
}

// Multiply 执行乘法运算
func (c *Calculator) Multiply(a, b int) int {
	result := a * b
	c.logger.LogFormat("Multiply(%d, %d) = %d", a, b, result)
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
	c.logger.LogFormat("Factorial(%d) = %d", n, result)
	return result
}

// MaxOfThree 返回三个数中的最大值，使用 common 包的 Max 函数
func (c *Calculator) MaxOfThree(a, b, num int) int {
	result := common.Max(common.Max(a, b), num)
	c.logger.LogFormat("MaxOfThree(%d, %d, %d) = %d", a, b, num, result)
	return result
}
