// 2129. 将标题首字母大写
// 给你一个字符串 title ，它由单个空格连接一个或多个单词组成，每个单词都只包含英文字母。请你按以下规则将每个单词的首字母 大写 ：

// 如果单词的长度为 1 或者 2 ，所有字母变成小写。
// 否则，将单词首字母大写，剩余字母变成小写。
// 请你返回 大写后 的 title 。

package main

import (
	"fmt"
	"strings"
)

func main() {
	title := "capiTalIze tHe titLe"
	title = "First leTTeR of EACH Word"
	// title = "Te"
	fmt.Println(capitalizeTitle(title))
}

func capitalizeTitle(title string) string {
	w := ""
	words := strings.Split(title, " ")
	if len(words) == 1 {
		if len(words[0]) <= 2 {
			return strings.ToLower(words[0])
		}
	}
	for i := 0; i < len(words); i++ {
		lower := strings.ToLower(words[i])
		if len(words[i]) <= 2 {
			w += lower
			if i != len(words)-1 {
				w += " "
			}
		} else {
			start := lower[0]
			startUpper := strings.ToUpper(string(start))
			w += startUpper + string(lower[1:])
			if i != len(words)-1 {
				w += " "
			}
		}
	}
	return w

}
