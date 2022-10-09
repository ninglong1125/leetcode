package main

import "fmt"

type Ring struct {
	Prev  *Ring
	Value interface{}
	Next  *Ring
}

func main() {
	r := New(5)
	fmt.Println(r)
}

func (r *Ring) initNode() {
	r.Next = r
	r.Prev = r

}

func New(n int) *Ring {
	if n < 1 {
		return nil
	}

	r := new(Ring)
	p := r
	for i := 0; i < n; i++ {
		p.Next = &Ring{Prev: p.Prev}
		p = p.Next
	}
	p.Next = r
	r.Prev = p
	return r
}
