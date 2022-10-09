// 给你一个字符串 text ，该字符串由若干被空格包围的单词组成。每个单词由一个或者多个小写英文字母组成，并且两个单词之间至少存在一个空格。题目测试用例保证 text 至少包含一个单词 。

// 请你重新排列空格，使每对相邻单词之间的空格数目都 相等 ，并尽可能 最大化 该数目。如果不能重新平均分配所有空格，请 将多余的空格放置在字符串末尾 ，
// 这也意味着返回的字符串应当与原 text 字符串的长度相等。

// 返回 重新排列空格后的字符串 。

// 输入：text = "  this   is  a sentence "
// 输出："this   is   a   sentence"

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "  this   is  a sentence "
	text = " practice   makes   perfect"
	text = "a"
	text = " hello"
	res := reorderSpaces(text)
	fmt.Println(res)
}

func reorderSpaces(text string) string {
	if len(text) < 2 {
		return text
	}
	result := strings.TrimSuffix(text, "")
	r := strings.Split(text, " ")
	newStr := ""
	count := 0
	words := 0
	for _, m := range result {
		if m == 32 {
			count++
		}
	}

	for _, m := range r {
		if len(m) != 0 {
			words++
		}
	}
	if words == 1 {
		for _, m := range r {
			if len(m) != 0 {
				newStr += m
				for j := 0; j < count; j++ {
					newStr += " "
				}
			}
		}
		return newStr
	}
	seg := count / (words - 1)
	for _, m := range r {
		if len(m) != 0 {
			newStr += m
			for j := 0; j < seg; j++ {
				newStr += " "
			}
		}
	}
	newStr = newStr[:len(newStr)-seg]
	mod := count % (words - 1)

	if mod != 0 {
		for i := 0; i < mod; i++ {
			newStr += " "
		}
	}

	return newStr
}
