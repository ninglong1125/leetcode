// 2413. 最小偶倍数
// 给你一个正整数 n ，返回 2 和 n 的最小公倍数（正整数）。

package main

import "fmt"

func main() {
	n := 16
	fmt.Println(smallestEvenMultiple(n))
}

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}
