// fmttest
package main

import (
	"fmt"
)

func main() {
	// Print 不换行
	fmt.Print("Print ... \n")
	// Println 换行
	fmt.Println("Hello World!")
	// Printf
	fmt.Printf("a = %d, b = %d, str = %s \n", 2, 3, "abc")

	str := fmt.Sprint("Sprint ... \n")
	fmt.Print(str)
}
