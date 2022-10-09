// 870. 优势洗牌
// 给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数目来描述。

// 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{2, 7, 11, 15, 6}
	nums2 := []int{1, 10, 4, 11, 6}
	fmt.Println(advantageCount(nums1, nums2))
}

func advantageCount(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	stack := make([]int, 0)
	length := len(nums1)
	sort.Ints(nums1)
	sort.Ints(nums2)
	fmt.Println(nums1)
	fmt.Println(nums2)
	i := 0
	j := 0
	for j < length && i < length {
		if nums1[i] > nums2[j] {
			result = append(result, nums1[i])
			i++
			j++

		} else {
			fmt.Println(nums1[i])
			stack = append(stack, nums1[i])
			i++

		}
	}
	fmt.Println(stack)
	result = append(result, stack...)
	return result
}

// func advantageCount(nums1 []int, nums2 []int) []int {
// 	n := len(nums1)
// 	idx1 := make([]int, n)
// 	idx2 := make([]int, n)
// 	for i := 1; i < n; i++ {
// 		idx1[i] = i
// 		idx2[i] = i
// 	}
// 	sort.Slice(idx1, func(i, j int) bool { return nums1[idx1[i]] < nums1[idx1[j]] })
// 	sort.Slice(idx2, func(i, j int) bool { return nums2[idx2[i]] < nums2[idx2[j]] })
// 	fmt.Println(idx1, idx2)
// 	ans := make([]int, n)
// 	left, right := 0, n-1
// 	for i := 0; i < n; i++ {
// 		if nums1[idx1[i]] > nums2[idx2[left]] {
// 			ans[idx2[left]] = nums1[idx1[i]]
// 			left++
// 		} else {
// 			ans[idx2[right]] = nums1[idx1[i]]
// 			right--
// 		}
// 	}
// 	return ans
// }
