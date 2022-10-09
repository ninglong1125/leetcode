// 剑指 Offer II 001. 整数除法
// 给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。

// 注意：

// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1

// 示例 1：

// 输入：a = 15, b = 2
// 输出：7
// 解释：15/2 = truncate(7.5) = 7

package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2147483647
	b := 1
	fmt.Println(divide(a, b))
}

func divide(a int, b int) int64 {
	if a == math.MinInt32 { // 考虑被除数为最小值的情况
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if a == math.MaxInt32 { // 考虑被除数为最小值的情况
		if b == 1 {
			return math.MaxInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if b == math.MinInt32 { // 考虑除数为最小值的情况
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}

	if a == 0 {
		return 0

	}
	if a < 0 || b < 0 {
		if a < 0 && b < 0 {
			a = -a
			b = -b
			return calc(a, b)
		} else if a < 0 && b > 0 {
			a = -a
			res := -1 * calc(a, b)
			if res < math.MinInt32 {
				return math.MinInt32
			}
			return res
		} else {
			b = -b
			res := -1 * calc(a, b)
			if res < math.MinInt32 {
				return math.MinInt32
			}
			return res
		}
	}
	return calc(a, b)
}

func calc(a, b int) int64 {
	result := int64(0)
	if a < b {
		return 0
	} else {
		fmt.Println(math.Pow(2.0, (math.Log2(float64(a)) - math.Log2(float64(b)))))
		result = int64(math.Pow(2.0, (math.Log2(float64(a)) - math.Log2(float64(b)))))
		fmt.Println(result)
		if result >= math.MaxInt32 {
			return math.MaxInt32
		}
		return result
	}
}
