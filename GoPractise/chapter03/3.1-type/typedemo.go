// typedemo
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

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	a.Add(2)
	fmt.Println("a = ", a)

	// 数组为纯粹的值类型
	var a1 = [3]int{1, 2, 3}
	var b1 = a1 // 数组内容的完整复制
	b1[1]++
	fmt.Println(a1, b1)

	var b2 = &a1 // 数组内容引用
	b2[1]++
	fmt.Println(a1, *b2)
}
