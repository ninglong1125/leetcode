package main

import (
	"fmt"
)

func main() {
	m := make(map[int]*int)
	s := []int{1, 2, 3}
	for _, v := range s {
		fmt.Println(v)
		m[v] = &v
	}
	fmt.Println(*m[1], *m[2], *m[3])
	fmt.Println(*m[2])
}
