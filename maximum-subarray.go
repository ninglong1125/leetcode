// 53. 最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 子数组 是数组中的一个连续部分。

// 示例 1：

// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums = []int{-1}
	nums = []int{-1, 5, -4}
	fmt.Println(maxSubArray(nums))
}

// 对于含有正数的序列而言,最大子序列肯定是正数,
// 所以头尾肯定都是正数.我们可以从第一个正数开始算起,
// 每往后加一个数便更新一次和的最大值;当当前和成为负数时,
// 则表明此前序列无法为后面提供最大子序列和,因此必须重新确定序列首项.
func maxSubArray(nums []int) int {
	length := len(nums)
	temp := make([]int, length)
	copy(temp, nums)
	sort.Ints(temp)
	if temp[length-1] < 0 {
		return temp[length-1]
	}
	maxsum := 0
	thissum := 0
	for i := 0; i < length; i++ {
		thissum += nums[i]

		if thissum > maxsum {
			maxsum = thissum
		} else if thissum < 0 {
			thissum = 0
		}
	}
	return maxsum
}
