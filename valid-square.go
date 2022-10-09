// 给定2D空间中四个点的坐标 p1, p2, p3 和 p4，如果这四个点构成一个正方形，则返回 true 。

// 点的坐标 pi 表示为 [xi, yi] 。输入 不是 按任何顺序给出的。

// 一个 有效的正方形 有四条等边和四个等角(90度角)。

package main

import (
	"fmt"
	"math"
)

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	p1p2 := calcDistance(p1, p2)
	p3p4 := calcDistance(p3, p4)
	p1p3 := calcDistance(p1, p3)
	p2p4 := calcDistance(p2, p4)
	p1p4 := calcDistance(p1, p4)
	p2p3 := calcDistance(p2, p3)
	if p1p2 != p3p4 || p1p3 != p2p4 || p1p4 != p2p3 {
		return false
	}
	if p1p2 == 0 || p3p4 == 0 || p1p3 == 0 {
		return false
	}
	if p1p2 != p1p3 {
		miner := math.Min(p1p2, p1p3)
		return miner == p1p4
	}

	return true
}

func calcDistance(x, y []int) float64 {

	return math.Sqrt(float64((y[1]-x[1])*(y[1]-x[1]) + (y[0]-x[0])*(y[0]-x[0])))
}
func main() {
	p1 := []int{0, 0}
	p2 := []int{5, 0}
	p3 := []int{5, 4}
	p4 := []int{0, 4}
	fmt.Println(validSquare(p1, p2, p3, p4))
}
