package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()

	}
	wg.Wait()
}
