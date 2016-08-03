// ex1
package main

import (
	"fmt"
)

func main() {
	r1 := defertest1()
	fmt.Println("r1 =", r1)

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()

	r2 := defertest2()
	fmt.Println("r2 =", r2)
}

func f() {
	fmt.Println("a")
	panic(55) // error message
	fmt.Println("b")
	fmt.Println("f")
}

// defer用来添加函数结束时执行的语句
func defertest1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func defertest2() (result int) {
	return 0
	defer func() {
		result++
	}()
	return 0
}
