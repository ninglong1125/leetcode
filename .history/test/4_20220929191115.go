package main

import (
	"fmt"
)

func main() {
	var (
		k = 4
		v = -1
		m = map[int]int{1: 1, 2: 2, 3: 3}
	)
	if v, hit := m[k]; hit {
		k += v
	}
	fmt.Println(k, v)
}
