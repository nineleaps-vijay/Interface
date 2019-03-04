package main

import (
	"fmt"
)

type Square struct {
	length  int
	breadth int
}

type Rectangle struct {
	length  int
	breadth int
}

type GetArea interface {
	Area() int
}

func (r *Rectangle) Area() int {
	return r.length * r.breadth
}

func (s *Square) Area() int {
	return s.length * s.breadth
}

func GenericAreaFunc(shapes ...GetArea) {
	for _, s := range shapes{
		fmt.Println(s.Area())
	}
}


func main() {
	s := Square{3, 3}
	r := Rectangle{4, 6}
	GenericAreaFunc(&s, &r)
	
}
