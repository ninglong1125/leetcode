// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	target = 13
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	ltemp := 0
	rtemp := len(nums) - 1
	for {
		if ltemp > rtemp {
			return -1
		}
		mid := (ltemp + rtemp) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			rtemp = mid - 1
		} else {
			ltemp = mid + 1
		}
	}
}
