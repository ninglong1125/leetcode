package main

import (
	"fmt"
)

func main() {
	var cnt int
	for i := 1; i <= 100; i++ {
		if i&1 == 0 {
			cnt += i
		} else {
			cnt -= i
		}
	}
	fmt.Println(cnt)
}
