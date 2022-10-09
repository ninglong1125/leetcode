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

func advantageCount(nums1, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	sort.Ints(nums1)
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool { return nums2[ids[i]] < nums2[ids[j]] })
	fmt.Println(ids)
	left, right := 0, n-1
	for _, x := range nums1 {
		if x > nums2[ids[left]] {
			ans[ids[left]] = x
			left++
		} else {
			ans[ids[right]] = x
			right--
		}
	}
	return ans
}
