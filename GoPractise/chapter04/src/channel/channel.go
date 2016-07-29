// channel
package main

import (
	"fmt"
)

func Count(ch chan int) {
	// 通过ch <- 1语句向对应的channel中写入一个数据
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	// 定义包含10个channel的数组
	chs := make([]chan int, 10)
	fmt.Println("==================")
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	fmt.Println("----------------")
	for _, ch := range chs {
		// 通过<-ch语句从10个channel中依次读取数据
		<-ch
	}

	// 调用make()时将缓冲区大小作为第二个参数传入
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
}
