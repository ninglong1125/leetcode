// 给你一个字符串数组 words ，数组中的每个字符串都可以看作是一个单词。请你按 任意 顺序返回 words 中是其他单词的子字符串的所有单词。

// 如果你可以删除 words[j] 最左侧和/或最右侧的若干字符得到 word[i] ，那么字符串 words[i] 就是 words[j] 的一个子字符串。

//

// 示例 1：

// 输入：words = ["mass","as","hero","superhero"]
// 输出：["as","hero"]
// 解释："as" 是 "mass" 的子字符串，"hero" 是 "superhero" 的子字符串。
// ["hero","as"] 也是有效的答案。
// 示例 2：

// 输入：words = ["leetcode","et","code"]
// 输出：["et","code"]
// 解释："et" 和 "code" 都是 "leetcode" 的子字符串。
// 示例 3：

// 输入：words = ["blue","green","bu"]
// 输出：[]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/string-matching-in-an-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	words2 := []string{"leetcode", "et", "code"}
	words3 := []string{"blue", "green", "bu"}
	words4 := []string{"leetcoder", "leetcode", "od", "hamlet", "am"}
	fmt.Println(stringMatching(words))
	fmt.Println(stringMatching(words2))
	fmt.Println(stringMatching(words3))
	fmt.Println(stringMatching(words4))
}

func stringMatching(words []string) []string {
	length := len(words)
	result := make([]string, 0)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j]) // 按字符串长度 升序 排列
	})
	for i := 0; i < length; i++ {
		str := words[i]
		flag := false

		for j := i + 1; j < length; j++ {
			if strings.Contains(words[j], str) {
				flag = true
			}
		}
		if flag {
			result = append(result, str)
			flag = false
		}
	}
	return result
}
