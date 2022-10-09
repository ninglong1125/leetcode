// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

// 给你一个整数数组  flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false。

//

// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/can-place-flowers
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

func main() {
	flowerbed := []int{1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0}
	flowerbed = []int{1, 0, 0, 0, 0, 1}
	n := 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	result := false
	if flowerbed[0] == 1 && len(flowerbed) == 1 {
		return false
	}
	if flowerbed[0] == 0 && len(flowerbed) == 1 {
		if n == 1 {
			return true
		}
		return false
	}
	count := 0
	flowerbed = append(flowerbed, 0)
	added := append([]int{0}, flowerbed...)
	for i := 1; i < len(added)-1; i++ {
		if added[i] == 0 && added[i-1] == 0 && added[i+1] == 0 {
			count += 1
		}
	}
	if count >= n {
		result = true
	}
	return result
}
