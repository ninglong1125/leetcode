package main

import (
	"fmt"
)

func main() {
	// var x interface{} = nil
	var phone Phone = new(Sangfor)
}

type Phone interface {
	Call(telephone *string) error
}
type Sangfor struct {
}

func (g Sangfor) Call(telephone *string) error {
	fmt.Println("sangfor call")
	return nil
}
