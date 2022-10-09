// 面试题 01.09. 字符串轮转
// 字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

// 示例1:

//  输入：s1 = "waterbottle", s2 = "erbottlewat"
//  输出：True

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	fmt.Println(isFlipedString(s1, s2))
}

func isFlipedString(s1 string, s2 string) bool {
	var str strings.Builder
	str.WriteString(s1)
	str.WriteString(s1)
	fmt.Println(str.String())
	if strings.Contains(str.String(), s2) {
		return true
	}
	return false
}
