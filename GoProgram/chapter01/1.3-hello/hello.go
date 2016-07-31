// hello
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "World!"

	// D:\GitHub\book\GoProgram\chapter01\1.3-hello\1.3-hello.exe
	fmt.Println("os.Args[0] =", os.Args[0])

	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}
