// 870. 优势洗牌
// 给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数目来描述。

// 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。

package main

import "sort"

func main() {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{1, 10, 4, 11}
	advantageCount(nums1, nums2)
}

func advantageCount(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	sort.Ints(nums1)

	return result
}
