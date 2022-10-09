// 有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。

// 示例 1:

// 输入: k = 5

// 输出: 9

// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/get-kth-magic-number-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "fmt"

func main() {
	n := 5
	fmt.Println(getKthMagicNumber(n))
}

func getKthMagicNumber(k int) int {
	dp := make([]int, k+1)
	dp[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i < k+1; i++ {
		x2, x3, x5 := dp[p3]*3, dp[p5]*5, dp[p7]*7
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p3++
		}
		if dp[i] == x3 {
			p5++
		}
		if dp[i] == x5 {
			p7++
		}
	}
	return dp[k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
