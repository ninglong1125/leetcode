// 1496. 判断路径是否相交
// 给你一个字符串 path，其中 path[i] 的值可以是 'N'、'S'、'E' 或者 'W'，分别表示向北、向南、向东、向西移动一个单位。

// 你从二维平面上的原点 (0, 0) 处开始出发，按 path 所指示的路径行走。

// 如果路径在任何位置上与自身相交，也就是走到之前已经走过的位置，请返回 true ；否则，返回 false 。

package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "NESWW"
	fmt.Println(isPathCrossing(path))
}
func isPathCrossing(path string) bool {
	mapping := make(map[string]int, 0)
	nums := []int{0, 0}
	flag := false
	strs := strings.Split(path, "")
	nstrs := fmt.Sprint(nums)

	mapping[nstrs] += 1
	for _, m := range strs {
		if m == "N" {
			nums[1] += 1

		} else if m == "E" {
			nums[0] += 1

		} else if m == "S" {
			nums[1] -= 1
		} else {
			nums[0] -= 1
		}
		nstrs := fmt.Sprint(nums)
		mapping[nstrs] += 1

	}

	for _, v := range mapping {
		if v > 1 {
			flag = true
		}
	}
	return flag
}
