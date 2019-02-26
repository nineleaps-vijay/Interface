package main

import (
	"fmt"
)

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first  string
	second string
}

func (s *StringPair) Exchange() {
	s.first, s.second = s.second, s.first
}

type Point [2]int

func (p *Point) Exchange() {
	p[0], p[1] = p[1], p[0]
}

func exchangeThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func main() {
	p := Point{1,2}
	fmt.Println(p)
	s := StringPair{"vijay", "shanker"}
	fmt.Println(s)
	exchangeThese(&p, &s)
	fmt.Println(p)
	fmt.Println(s)	
}