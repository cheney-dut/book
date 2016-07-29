// interface
package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

func main() {
	var a Integer = 1
	var b LessAdder = &a

	fmt.Println(b.Less(2))
	b.Add(2)
	fmt.Println(b)

	// 异常
	// var b1 LessAdder = a
	// fmt.Println(b1.Less(2))

	var c11 Lesser = &a
	var c12 Lesser = a
	fmt.Println(c11.Less(2))
	fmt.Println(c12.Less(2))
}
