// 1816. 截断句子
// 句子 是一个单词列表，列表中的单词之间用单个空格隔开，且不存在前导或尾随空格。每个单词仅由大小写英文字母组成（不含标点符号）。

// 例如，"Hello World"、"HELLO" 和 "hello world hello world" 都是句子。
// 给你一个句子 s​​​​​​ 和一个整数 k​​​​​​ ，请你将 s​​ 截断 ​，​​​使截断后的句子仅含 前 k​​​​​​ 个单词。返回 截断 s​​​​​​ 后得到的句子。

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello how are you  Contestant"
	k := 4
	fmt.Println(truncateSentence(s, k))
}

func truncateSentence(s string, k int) string {
	var result strings.Builder
	res := strings.Split(s, " ")
	for i := 0; i < k; i++ {
		if len(res[i]) == 0 {
			continue
		}
		result.WriteString(res[i])
		result.WriteString(" ")
	}
	return result.String()[:result.Len()-1]
}
