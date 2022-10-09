// 1492. n 的第 k 个因子
// 给你两个正整数 n 和 k 。

// 如果正整数 i 满足 n % i == 0 ，那么我们就说正整数 i 是整数 n 的因子。

// 考虑整数 n 的所有因子，将它们 升序排列 。请你返回第 k 个因子。如果 n 的因子数少于 k ，请你返回 -1 。

// 示例 1：

// 输入：n = 12, k = 3
// 输出：3
// 解释：因子列表包括 [1, 2, 3, 4, 6, 12]，第 3 个因子是 3 。
package main

import "fmt"

func main() {
	n := 12
	k := 3
	n = 4
	k = 4
	n = 7
	k = 2
	fmt.Println(kthFactor(n, k))

}

func kthFactor(n int, k int) int {
	results := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			results = append(results, i)
		}
	}
	if k > len(results) {
		return -1
	}
	return results[k-1]
}
