// 856. 括号的分数
// 给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：

// () 得 1 分。
// AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
// (A) 得 2 * A 分，其中 A 是平衡括号字符串。

package main

import "fmt"

func main() {
	s := "()"
	fmt.Println(scoreOfParentheses(s))
}

func scoreOfParentheses(s string) int {
	if len(s) == 2 {
		return 1
	}
	result := 0

	return result
}
