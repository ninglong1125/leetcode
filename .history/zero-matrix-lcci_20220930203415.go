// 面试题 01.08. 零矩阵
// 编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

// 示例 1：

// 输入：
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// 输出：
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]
package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{
			1, 1, 1,
		},
		[]int{
			1, 0, 1,
		},
		[]int{
			1, 1, 1,
		},
	}
	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	mcopy := make([][]int, len(matrix))
	copy(mcopy, matrix)
	fmt.Println(mcopy)
}
