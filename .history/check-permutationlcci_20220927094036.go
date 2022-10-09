// 面试题 01.02. 判定是否互为字符重排
// 给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

// 示例 1：

// 输入: s1 = "abc", s2 = "bca"
// 输出: true

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "abc"
	s2 := "cba"
	fmt.Println(CheckPermutation(s1, s2))
}

func CheckPermutation(s1 string, s2 string) bool {
	r1 := strings.Split(s1, "")
	r2 := strings.Split(s2, "")
	sort.Strings(r1)
	sort.Strings(r2)
	ss1, ss2 := "", ""
	for _, m := range r1 {
		ss1 += m
	}
	for _, m := range r2 {
		ss2 += m
	}

	return ss1 == ss2
}
