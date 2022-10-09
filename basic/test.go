package main

import "fmt"

func main() {
	ch := make(chan struct{})
	defer close(ch)

	go func() {
		ch <- struct{}{}
	}()

	i := 0
	go func() {
		for range ch {
			i++
			<-ch
		}
	}()

	fmt.Println(i)
}
