// 409. 最长回文串
// 给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串 。

// 在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。

package main

import "fmt"

func main() {
	s := "Aa"
	s = "ccc"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(s string) int {
	result := 0
	counter := make(map[string]int, 0)
	if len(s) == 1 {
		return 1
	}
	for _, m := range s {
		counter[string(m)] += 1
	}

	for _, v := range counter {
		if v%2 == 0 {
			result += v
		} else if v%2 == 1 {
			result += v - 1

		}
	}

	return result
}
