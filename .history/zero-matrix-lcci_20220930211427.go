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
			0, 1, 2, 1,
		},
		[]int{
			3, 4, 5, 2,
		},
		[]int{
			1, 3, 1, 5,
		},
	}
	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	M := len(matrix)
	N := len(matrix[0])
	mcopy := make([][]int, 0)
	for i := 0; i < M; i++ {
		m := make([]int, 0)
		for j := 0; j < N; j++ {
			m = append(m, matrix[i][j])
		}
		mcopy = append(mcopy, m)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < len(matrix[i]); k++ {
					mcopy[i][k] = 0
				}

			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				fmt.Println(i, j)
				for m := 0; m < len(matrix); m++ {
					matrix[m][j] = 0
				}
			}
		}
	}
	fmt.Println(mcopy, matrix)
	for i := 0; i < len(mcopy); i++ {
		for j := 0; j < len(mcopy[i]); j++ {
			matrix[i][j] = matrix[i][j] & mcopy[i][j]
		}
	}
	fmt.Println(matrix)
}
