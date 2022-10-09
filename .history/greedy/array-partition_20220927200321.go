// 561. 数组拆分
// 给定长度为 2n 的整数数组 nums ，你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从 1 到 n 的 min(ai, bi) 总和最大。

// 返回该 最大总和 。

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 1, 3, 4}
	nums = []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	result := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		min := compareMin(nums[i], nums[i+1])
		result += min
	}

	return result
}

func compareMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
