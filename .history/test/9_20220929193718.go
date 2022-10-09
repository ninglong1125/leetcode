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

type Counter struct {
	Adder
}

func main() {
	var (
		a int32 = 3
		b int32 = 2
		c       = Counter{Adder: Adder{A: &a, B: b}}
	)
	for i := 0; i < 2; i++ {
		c.Add(int32(i))
	}
	fmt.Println(*c.A + c.B + a + b)
}
