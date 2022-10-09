// 856. 括号的分数
// 给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：

// () 得 1 分。
// AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
// (A) 得 2 * A 分，其中 A 是平衡括号字符串。

package main

import "fmt"

func main() {
	s := "()()"
	s = "(()(()))"
	fmt.Println(scoreOfParentheses(s))
}

func scoreOfParentheses(s string) int {
	stack := make([]int, 0)
	if len(s) == 2 {
		return 1
	}
	stack = append(stack, 0)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack = append(stack, 0)
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[0] += max(2*last, 1)
		}
	}
	return stack[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
