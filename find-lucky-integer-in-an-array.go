// 1394. 找出数组中的幸运数
// 在整数数组中，如果一个整数的出现频次和它的数值大小相等，我们就称这个整数为「幸运数」。

// 给你一个整数数组 arr，请你从中找出并返回一个幸运数。

// 如果数组中存在多个幸运数，只需返回 最大 的那个。
// 如果数组中不含幸运数，则返回 -1 。

// 示例 1：

package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 2, 2, 3, 3}
	arr = []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println(findLucky(arr))
}

func findLucky(arr []int) int {
	if len(arr) == 1 {
		return -1
	}
	counter := make(map[int]int, 0)
	for _, m := range arr {
		counter[m] += 1
	}
	max := -2
	for k, v := range counter {
		if k == v && k > max {
			max = k
		}
	}

	if max == -2 {
		return -1
	}
	return max
}
