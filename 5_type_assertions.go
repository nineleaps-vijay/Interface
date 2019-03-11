package main

import (
	"fmt"
	"reflect"
)

type Namer interface {
	getName() string
}

type Person struct {
	firstName string
	lastName  string
}

func (p Person) getName() string {
	return p.firstName + " " + p.lastName
}

func PrintName(n Namer) string {
	return n.getName()
}

func main() {
	p := Person{"vijay", "shanker"}
	x := PrintName(p)
	fmt.Println(x)

	var v Namer
	fmt.Println(reflect.TypeOf(v))
	v = Person{"Ut", "kr"}

	foo, _ := v.(Namer)
	fmt.Println(foo)
	
	//bar := v.(string)

	fmt.Println(reflect.TypeOf(foo))
	var w Namer
	switch w := v.(type) {
	case nil:
		fmt.Println("abc is nil")
	case Person:
		fmt.Println("abc is Person struct")
	default:
		fmt.Println('abc')

	}

}
