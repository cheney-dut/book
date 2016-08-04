// str1
// 参考：
// go语言 看代码，学strings包：http://blog.csdn.net/zistxym/article/details/17582265
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(" Contains 函数的用法")
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true 这里要特别注意
	fmt.Println(strings.Contains("我是中国人", "我"))     //true
	fmt.Println("--------------------------------")

	fmt.Println(" ContainsAny 函数的用法")
	fmt.Println(strings.ContainsAny("team", "i"))         // false
	fmt.Println(strings.ContainsAny("failure", "u & i"))  // true
	fmt.Println(strings.ContainsAny("failure", "t || i")) // true
	fmt.Println(strings.ContainsAny("foo", ""))           // false
	fmt.Println(strings.ContainsAny("", ""))              // false
	fmt.Println("--------------------------------")

	fmt.Println(" CountainsRune 函数的用法")
	fmt.Println(strings.ContainsRune("我是中国", '我')) // true 注意第二个参数，用的是字符
	fmt.Println("--------------------------------")

	fmt.Println(" Count 函数的用法")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune result: 5，源码中有实现
	fmt.Println("--------------------------------")

	fmt.Println(" EqualFold 函数的用法")
	fmt.Println(strings.EqualFold("Go", "go")) // 大小写忽略
	fmt.Println("--------------------------------")

	// 整数占位符
	// %q 单引号围绕的字符字面值，由Go语法安全地转义 Printf("%q", 0x4E2D) '中'
	// 字符串与字节切片
	// %q 双引号围绕的字符串，由Go语法安全地转义 Printf("%q", "Go语言") "Go语言"
	fmt.Println(" Fields 函数的用法")
	fmt.Println("Fields are:", strings.Fields(" foo bar baz "))     // Fields are: [foo bar baz]
	fmt.Printf("Fields are: %q\n", strings.Fields(" foo bar baz ")) // Fields are: ["foo" "bar" "baz"] 返回一个列表
	fmt.Println("")

	//相当于用函数做为参数，支持匿名函数
	fmt.Println(" FieldsFunc 函数的用法")
	for _, record := range []string{" aaa*1892*122", "aaa\taa\t", "124|939|22"} {
		fmt.Println(strings.FieldsFunc(record, func(ch rune) bool {
			switch {
			case ch > '5': // TODO
				return true
			}
			return false
		}))
	}
	fmt.Println("--------------------------------")

	fmt.Println(" HasPrefix 函数的用法")                  // 类似Java中的startWith
	fmt.Println(strings.HasPrefix("NLT_abc", "NLT")) // 前缀是以NLT开头的
	fmt.Println(strings.HasPrefix("NLT_abc", "nlt")) // 区分大小写
	fmt.Println("--------------------------------")

	fmt.Println(" HasSuffix 函数的用法")                  // 类似Java中的endWith
	fmt.Println(strings.HasSuffix("NTL_abc", "abc")) // 后缀是以abc开头的
	fmt.Println(strings.HasSuffix("NTL_abc", "Abc"))
	fmt.Println("--------------------------------")

	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.Index("NLT_abc", "abc")) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.Index("NLT_abc", "aaa")) // 不存在返回 -1
	fmt.Println(strings.Index("我是中国人", "中"))     // 存在返回 6，中文字符是用3个字节存的
	fmt.Println("--------------------------------")

	fmt.Println(" IndexAny 函数的用法") // ContainsAny源码中使用的就是IndexAny(s, chars) >= 0
	fmt.Println("--------------------------------")

	fmt.Println(" Join 函数的用法")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // 返回字符串：foo, bar, baz
	fmt.Println("--------------------------------")

	fmt.Println(" Map 函数的用法") // 自定义映射关系进行转换
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	fmt.Println("--------------------------------")

	fmt.Println(" Repeat 函数的用法")
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana
	fmt.Println("--------------------------------")

	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 4))      // oinky oinky oinky
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
	fmt.Println("--------------------------------")

	fmt.Println(" Split 函数的用法")
	fmt.Println(strings.Split("a,b,c", ","))                               // [a b c]
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]
	fmt.Println("--------------------------------")

	fmt.Println(" SplitAfter 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfter("/home/m_ta/src", "/")) //["/" "home/" "m_ta/" "src"]
	fmt.Println("--------------------------------")

	// n > 0: at most n substrings; the last substring will be the unsplit remainder.
	// n == 0: the result is nil (zero substrings)
	// n < 0: all substrings
	fmt.Println(" SplitAfterN 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfterN("/home/m_ta/src", "/", 2))  //["/" "home/m_ta/src"]
	fmt.Printf("%q\n", strings.SplitAfterN("#home#m_ta#src", "#", -1)) //["#" "home#" "m_ta#" "src"]
	fmt.Println("--------------------------------")

	fmt.Println(" SplitN 函数的用法")
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 1)) // ["/home/m_ta/src"]

	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 2))  // ["" "home/m_ta/src"]
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", -1)) // ["" "home" "m_ta" "src"]
	fmt.Printf("%q\n", strings.SplitN("home,m_ta,src", ",", 2))   // ["home" "m_ta,src"]

	fmt.Printf("%q\n", strings.SplitN("#home#m_ta#src", "#", -1)) // ["" "home" "m_ta" "src"]
	fmt.Println("--------------------------------")

	fmt.Println(" ToTitle 函数的用法")
	fmt.Println(strings.ToTitle("loud noises")) // LOUD NOISES
	fmt.Println(strings.ToTitle("loud 中国"))     // LOUD 中国
	fmt.Println("--------------------------------")

	fmt.Println(" Trim  函数的用法")
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! ")) // ["Achtung"]
	fmt.Println(" TrimSpace 函数的用法")
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n ")) // a lone gopher
}
