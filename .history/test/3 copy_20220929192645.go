package main

import (
	"fmt"
)

func main() {
	var (
		m     = map[int]int{0: 3, 1: 2, 3: 0}
		s     = [4]int{3, 2, 1}
		count int
	)
	for _, v := range s {
		count += m[v]
	}
	fmt.Println(count)
}
