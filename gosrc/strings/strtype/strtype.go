// strtype.go
// 扩展strings，增加append方法
package main

import (
	"fmt"
	"strings"
)

type StrType string

func (str StrType) Append(s string) string {
	a := make([]string, 2)
	a[0] = string(str)
	a[1] = s
	return strings.Join(a, "")
}

func main() {
	var str StrType = "Hello"
	fmt.Println(str.Append(" World"))
}
