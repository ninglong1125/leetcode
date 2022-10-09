package main

import (
	"fmt"
	"sync/atomic"
)

type Adder struct {
	A *int32
	B int32
}

func (a Adder) Add(c int32) {
	a.B = atomic.AddInt32(a.A, c)
}

func main() {
	fmt.Println(a, b, c, d)
}
