// 给你一个字符串数组 words 和一个字符串 pref 。

// 返回 words 中以 pref 作为 前缀 的字符串的数目。

// 字符串 s 的 前缀 就是  s 的任一前导连续字符串。

//
package main

import (
	"fmt"
)

func main() {
	words := []string{
		"pay", "attention", "practice", "attend",
	}
	words = []string{
		"leetcode", "win", "loops", "success",
	}
	pref := "l"

	fmt.Println(prefixCount(words, pref))
}

func prefixCount(words []string, pref string) int {
	count := 0
	for _, s := range words {
		if len(s) < len(pref) {
			continue
		}
		flag := true
		for i, m := range pref {
			if string(s[i]) != string(m) {
				flag = false
				break
			}

		}
		if flag {
			count += 1
		}
	}

	return count
}

// func prefixCount(words []string, pref string) int {
// 	count := len(words)
// 	for _, s := range words {
// 		if len(s) < len(pref) {
// 			count--
// 			break
// 		}
// 		for i, m := range pref {
// 			if string(s[i]) != string(m) {
// 				count--
// 				break
// 			}

// 		}
// 	}

// 	return count
// }
