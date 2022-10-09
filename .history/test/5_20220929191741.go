package main

import (
	"fmt"
)

func main() {
	fmt.Println(demo(1, 2))
}

func demo(a, b int) (c int) {
	fmt.Println(c)
	defer func() {
		c++
	}()
	c++
	return a + b
}
