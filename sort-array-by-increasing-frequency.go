// 1636. 按照频率将数组升序排序
// 给你一个整数数组 nums ，请你将数组按照每个值的频率 升序 排序。如果有多个值的频率相同，请你按照数值本身将它们 降序 排序。

// 请你返回排序后的数组。

// 输入：nums = [1,1,2,2,2,3]
// 输出：[3,1,1,2,2,2]
// 解释：'3' 频率为 1，'1' 频率为 2，'2' 频率为 3 。
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2, 2, 3}
	nums = []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	nums = []int{5}
	fmt.Println(frequencySort(nums))
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (a PairList) Len() int      { return len(a) }
func (a PairList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PairList) Less(i, j int) bool {
	if a[i].Value == a[j].Value {
		return a[i].Key > a[j].Key
	}
	return a[i].Value < a[j].Value
}

func frequencySort(nums []int) []int {
	length := len(nums)
	if length == 1 {
		return nums
	}
	result := make([]int, 0)
	counter := make(map[int]int, 0)

	for i := 0; i < length; i++ {
		counter[nums[i]] += 1
	}

	p := make(PairList, 0)
	for k, v := range counter {
		p = append(p, Pair{k, v})
	}
	sort.Sort(p)
	for i := 0; i < p.Len(); i++ {
		item := p[i]
		for i := 0; i < item.Value; i++ {
			result = append(result, item.Key)
		}
	}
	return result
}
