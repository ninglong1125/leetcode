package main

import (
	"fmt"
)

func main() {
	s := "abc"
}

func defer_call() {
	defer func() {
		fmt.Println("a")
	}()

	defer func() {
		fmt.Println("b")

	}()
	panic("c")
	defer func() {
		fmt.Println("d")
	}()
}
