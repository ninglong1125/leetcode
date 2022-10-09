// 1684. 统计一致字符串的数目
// 给你一个由不同字符组成的字符串 allowed 和一个字符串数组 words 。如果一个字符串的每一个字符都在 allowed 中，就称这个字符串是 一致字符串 。

// 请你返回 words 数组中 一致字符串 的数目。

package main

import "fmt"

func main() {
	allowed := "ab"
	words := []string{
		"ad", "bd", "aaab", "baa", "badab",
	}
	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		word := set(words[i])
		fmt.Println(word)
	}

	return count
}

func set(str string) string {

	x := make(map[string]int, len(str))
	for i := 0; i < len(str); i++ {

		x[fmt.Sprintf("%c", str[i])] = i
	}
	i := ""
	for k, _ := range x {

		i += fmt.Sprint(k)
	}

	return i
}
