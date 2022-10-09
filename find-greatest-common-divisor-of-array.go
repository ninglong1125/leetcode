// 1979. 找出数组的最大公约数
// 给你一个整数数组 nums ，返回数组中最大数和最小数的 最大公约数 。

// 两个数的 最大公约数 是能够被两个数整除的最大正整数。

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 6, 9}
	fmt.Println(findGCD(nums))
}

func findGCD(nums []int) int {
	sort.Ints(nums)
	min, max := nums[0], nums[len(nums)-1]
	fmt.Println(min, max)
	res := gcd(max, min)
	return res
}

func gcd(max, min int) int {
	if min == 0 {
		return max
	}
	r := max % min
	return gcd(min, r)
}
