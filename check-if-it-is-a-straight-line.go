// 缀点成线
// 给定一个数组 coordinates ，其中 coordinates[i] = [x, y] ， [x, y] 表示横坐标为 x、纵坐标为 y 的点。请你来判断，这些点是否在该坐标系中属于同一条直线上。

// 输入：coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
// 输出：true

package main

import "fmt"

func main() {
	coordiantes := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7},
	}
	coordiantes2 := [][]int{
		{1, 2}, {1, 3}, {4, 3},
	}
	coordiantes3 := [][]int{
		{3, 3}, {3, 3},
	}
	coordiantes4 := [][]int{
		{-4, -3}, {1, 0}, {3, -1}, {0, -1}, {-5, 2},
	}
	fmt.Println(checkStraightLine(coordiantes))
	fmt.Println(checkStraightLine(coordiantes2))
	fmt.Println(checkStraightLine(coordiantes3))
	fmt.Println(checkStraightLine(coordiantes4))
}

func checkStraightLine(coordinates [][]int) bool {
	length := len(coordinates)
	flag := true
	if length <= 2 {
		return true
	}
	for i := 0; i < length-1; i++ {
		if (coordinates[i+1][0] - coordinates[i][0]) != 0 {
			x, y := coordinates[0], coordinates[1]
			if y[0]-x[0] == 0 {
				continue
			}
			k1 := float64(y[1]-x[1]) / float64(y[0]-x[0])
			k2 := float64(coordinates[i+1][1]-coordinates[i][1]) / float64(coordinates[i+1][0]-coordinates[i][0])
			if k1 != k2 {
				flag = false
				return flag
			}
		} else {
			for j := 0; j < length-1; j++ {
				if coordinates[j+1][0] != coordinates[j][0] {
					flag = false
					break
				}
			}
		}
	}

	return flag
}
