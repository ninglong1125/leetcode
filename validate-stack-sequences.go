// 给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，
// 只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。

// 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// 输出：true

package main

import "fmt"

func main() {
	pushed := []int{1, 2, 3}
	popped := []int{3, 2, 1}
	res := validateStackSequences(pushed, popped)
	fmt.Println(res)
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	length := len(pushed)
	j := 0
	for i := 0; i < length; i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}
