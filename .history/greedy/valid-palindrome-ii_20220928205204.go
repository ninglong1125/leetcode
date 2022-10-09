// 680. 验证回文串 II
// 给你一个字符串 s，最多 可以从中删除一个字符。

// 请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：s = "aba"
// 输出：true
package main

import "fmt"

func main() {
	s := "abba"
	fmt.Println(validPalindrome(s))
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			flag1, flag2 := true, true
			for i, j := i, j-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := i+1, j; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return false
}
