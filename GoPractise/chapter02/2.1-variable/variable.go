// variable
package main

import (
	"fmt"
)

func main() {
	//	var v1 int
	//	var v2 string
	var (
		v1 int = 10
		v2 string
	)
	v1 = 123
	var v12 = 12
	v13 := 13

	v1, v12 = v12, v1
	fmt.Println("int  v1: ", v1)
	fmt.Println("int v12: ", v12)
	fmt.Println("int v13: ", v13)
	fmt.Println("string: ", v2)

	// 数组
	var v3 [10]int
	fmt.Println("[10]int: ", v3)

	// 数组切片
	var v4 []int
	fmt.Println("[]int: ", v4)

	var v5 struct {
		f int
	}
	fmt.Println("struct: ", v5)

	// 指针
	var v6 *int
	fmt.Println("*int: ", v6)

	// map，key为string类型，value为int类型
	var v7 map[string]int
	fmt.Println("map[string]int: ", v7)

	var v8 func(a int) int
	//fmt.Println("func(a int) int: " + v8)
	fmt.Println(v8)
}
