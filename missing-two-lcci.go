// 面试题 17.19. 消失的两个数字
// 给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？

// 以任意顺序返回这两个数字均可。

// 示例 1:

// 输入: [1]
// 输出: [2,3]

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5, 6, 8}
	fmt.Println(missingTwo((nums)))
}

// func missingTwo(nums []int) []int {
// 	length := len(nums)
// 	result := make([]int, 0)
// 	counter := make(map[int]int, 0)
// 	for i := 0; i < length+2; i++ {
// 		counter[i+1] = 0
// 	}
// 	for _, m := range nums {
// 		counter[m] = 1
// 	}
// 	for k, v := range counter {
// 		if v == 0 {
// 			result = append(result, k)
// 		}
// 	}
// 	return result
// }

func missingTwo(nums []int) []int {
	result := make([]int, 0)
	length := len(nums)
	sums := ((1 + length + 2) * (length + 2)) / 2
	nsum := 0
	for i := 0; i < length; i++ {
		nsum += nums[i]
	}
	extra := sums - nsum
	divide := extra / 2

	dleftsum := ((1 + extra/2) * (extra / 2)) / 2
	for _, m := range nums {
		if m <= divide {
			dleftsum -= m
		}
	}
	result = append(result, dleftsum)
	result = append(result, extra-dleftsum)
	return result
}
