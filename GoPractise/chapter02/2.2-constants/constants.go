// constants.go
package main

import (
	"fmt"
)

func main() {
	const PI float64 = 3.14159265358979323846
	const zero = 0.0 // 无类型浮点常量
	const (
		siz int64 = 1024
		eof       = -1 // 无类型整型常量
	)
	// u = 0.0, v = 3.0， 常量的多重赋值
	const u, v float32 = 0, 3
	// 无类型整型和字符串常量
	const a, b, c = 3, 4, "foot"

	const mask = 1 << 3

	fmt.Println("PI = ", PI)
	fmt.Println("zero = ", zero)
	fmt.Println("siz = ", siz)
	fmt.Println("eof = ", eof)
	fmt.Println("u = ", u)
	fmt.Println("v = ", v)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("mask = ", mask)

	fmt.Println("-------------------------------")

	fmt.Println("预定义常量")
	// iota 被重设为0
	const (
		c0 = iota // 0
		c1 = iota // 1
		c2 = iota // 2
	)
	fmt.Println("c0 = ", c0)
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)

	const (
		a2 = 1 << iota // 1，（iota在每个const开头被重设为0）
		b2 = 1 << iota // 2
		d2 = 1 << iota // 4
	)
	fmt.Println("a2 = ", a2)
	fmt.Println("b2 = ", b2)
	fmt.Println("d2 = ", d2)

	const (
		u2         = iota * 42 // 0
		v2 float64 = iota * 42 // 42
		w2         = iota * 42 // 84
	)
	fmt.Println("u2 = ", u2)
	fmt.Println("v2 = ", v2)
	fmt.Println("w2 = ", w2)

	const x2 = iota // x == 0
	const y2 = iota // y == 0
	fmt.Println("x2 = ", x2)
	fmt.Println("y2 = ", y2)

	fmt.Println("-------------------------------")

	fmt.Println("枚举")
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出
	)
	fmt.Println("Sunday = ", Sunday)
	fmt.Println("Monday = ", Monday)
	fmt.Println("Tuesday = ", Tuesday)
	fmt.Println("Wednesday = ", Wednesday)
	fmt.Println("Thursday = ", Thursday)
	fmt.Println("Friday = ", Friday)
	fmt.Println("Saturday = ", Saturday)
	fmt.Println("numberOfDays = ", numberOfDays)
}
