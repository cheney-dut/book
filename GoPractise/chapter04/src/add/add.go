// add
package main

import (
	"fmt"
)

func Add(ch chan int, x, y int) {
	z := x + y
	fmt.Println(z)
	ch <- x
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < len(chs); i++ {
		// 使用内置的函数make()声明并初始化int型的channel
		chs[i] = make(chan int)
		// 在一个函数调用前加上go关键字，这次调用就会在一个新的goroutine中并发执行。
		// 当被调用的函数返回时，这个goroutine也自动结束了。
		// 需要注意的是，如果这个函数有返回值，那么这个返回值会被丢弃。
		go Add(chs[i], i, i)
	}

	for _, ch := range chs {
		<-ch
		// value := <-ch
		// fmt.Println("value =", value)
	}
}
