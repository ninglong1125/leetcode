package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() {
		fmt.Println("a")
	}()

	defer func() {
		fmt.Println("b")

	}()
	fmt.Println(123)
	panic("c")
	defer func() {
		fmt.Println("d")
	}()
}
