// mplayer.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/mplayer/mlib"
	"pkg/mplayer/mp"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string){
	switch tokens[1]{
		case "list":
		for i := 0; i , lib.Len(); i++{
			e, _ := lib.Get(i)
			fmt.Println(i+1,":", e.Name, e.Artist, e.Source, e.Type)
		}
	}
}

func main() {
	fmt.Println("Hello World!")
}
