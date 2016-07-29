// flow
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ifFlow() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := r.Intn(10)
	if a < 5 {
		fmt.Printf("a[%d] < 5", a)
	} else {
		fmt.Printf("a[%d] >= 5", a)
	}
	fmt.Println("\n---------------------")
}

func switchFlow() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := r.Intn(10)

	fmt.Printf("switch a[%d]: \n", a)
	switch a {
	case 0:
		fmt.Printf("a[%d] = 0 \n", a)
	case 1:
		fmt.Printf("a[%d] = 1 \n", a)
	case 2:
		fallthrough
	case 3:
		fmt.Printf("a[%d] = 3 \n", a)
	case 4, 5, 6:
		fmt.Printf("a[%d] in (4, 5, 6)", a)
	default:
		fmt.Println("Default")
	}
	fmt.Println("---------------------")
}

func forFlow() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
	fmt.Println("---------------------")
}

func main() {
	ifFlow()

	switchFlow()

	forFlow()
}
