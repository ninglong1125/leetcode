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
	mcopy2 := make([][]int, len(matrix))
	copy(mcopy, matrix)
	copy(mcopy2, matrix)
	fmt.Println(mcopy, mcopy2)
	for i := 0; i < len(mcopy); i++ {
		for j := 0; j < len(mcopy[i]); j++ {
			if mcopy[i][j] == 0 {
				for k := 0; k < len(mcopy[i]); k++ {
					mcopy2[i][k] = 0
				}

			}
			// fmt.Println(mcopy[i][j])
		}
	}
	fmt.Println(mcopy2)
	for i := 0; i < len(mcopy); i++ {
		for j := 0; j < len(mcopy[i]); j++ {
			if mcopy[i][j] == 0 {
				for m := 0; m < len(mcopy); m++ {
					matrix[m][j] = 0
				}
			}
			// fmt.Println(mcopy[i][j])
		}
	}
	fmt.Println(mcopy2)
	for i := 0; i < len(mcopy); i++ {
		for j := 0; j < len(mcopy[i]); j++ {
			mcopy[i][j] = matrix[i][j] & mcopy2[i][j]
		}
	}
	fmt.Println(5 & 5)
	fmt.Println(matrix)
}
