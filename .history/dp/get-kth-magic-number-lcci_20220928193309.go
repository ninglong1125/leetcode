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
	p1, p2, p3 := 1, 1, 1

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}