// 文件夹操作日志搜集器
// 每当用户执行变更文件夹操作时，LeetCode 文件系统都会保存一条日志记录。

// 下面给出对变更操作的说明：

// "../" ：移动到当前文件夹的父文件夹。如果已经在主文件夹下，则 继续停留在当前文件夹 。
// "./" ：继续停留在当前文件夹。
// "x/" ：移动到名为 x 的子文件夹中。题目数据 保证总是存在文件夹 x 。
// 给你一个字符串列表 logs ，其中 logs[i] 是用户在 ith 步执行的操作。

// 文件系统启动时位于主文件夹，然后执行 logs 中的操作。

// 执行完所有变更文件夹操作后，请你找出 返回主文件夹所需的最小步数 。

package main

import "fmt"

func main() {
	logs := []string{
		"d1/", "d2/", "../", "d21/", "./",
	}
	logs = []string{
		"d1/", "d2/", "./", "d3/", "../", "d31/",
	}
	logs = []string{
		"d1/", "../", "../", "../",
	}
	fmt.Println(minOperations(logs))
}

func minOperations(logs []string) int {
	depth := 0
	for _, m := range logs {
		if m == "../" {
			if depth > 0 {
				depth -= 1
			}
		} else if m == "./" {
			continue
		} else {
			depth += 1
		}

	}
	if depth < 0 {
		return 0
	}
	return depth
}

// func minOperations(logs []string) int {
// 	result := make([]string, 0)
// 	for i, m := range logs {
// 		if m == "../" {
// 			if len(result) == 0 {
// 				continue
// 			}
// 			result = result[:len(result)-1]

// 		} else if m == "./" {
// 			continue
// 		} else {
// 			result = append(result, logs[i])
// 		}

// 	}
// 	return len(result)
// }
