// 给你一个整数数组 arr ，请你将数组中的每个元素替换为它们排序后的序号。

// 序号代表了一个元素有多大。序号编号的规则如下：

// 序号从 1 开始编号。
// 一个元素越大，那么序号越大。如果两个元素相等，那么它们的序号相同。
// 每个数字的序号都应该尽可能地小。
//
package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{
		37, 12, 28, 9, 100, 56, 80, 5, 12,
		// 100, 100, 100,
	}
	arrayRankTransform(array)
}

func arrayRankTransform(arr []int) []int {
	dict := make(map[int]int, 0)
	array := make([]int, len(arr))
	copy(array, arr)
	fmt.Println(array)
	sort.Ints(array)
	for _, m := range array {
		if dict[m] != 0 {
			continue
		}
		dict[m] = len(dict) + 1
	}

	result := make([]int, 0)
	for _, m := range arr {
		result = append(result, dict[m])
	}
	fmt.Println(result)
	return result
}
