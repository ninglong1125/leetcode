package main

import (
	"fmt"
)

func main() {
	fmt.Println(demo(1, 2))
}

func demo(a, b int) (c int) {
	defer func() {
		c++
	}()
	c++
	return a + b
}
